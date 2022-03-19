from flask import Flask, request, jsonify, abort
from flask.app import HTTPException
from sqlalchemy import create_engine, Column, Integer
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.dialects.postgresql import JSONB
from sqlalchemy.orm import sessionmaker
import pandas as pd
import numpy as np
from io import StringIO
from sklearn.model_selection import KFold
from sklearn.linear_model import Ridge
import json
import logging
import sys

logging.basicConfig(stream=sys.stdout, format='%(asctime)s %(message)s', level=logging.DEBUG)


def mse_score(y_true, y_pred):
    result = np.mean((y_true - y_pred) ** 2)
    return result


# trainig function
def model_trainig(data, target, n_folds, fit_intercept, coefs):
    c = StringIO(data)
    try:
        df = pd.read_csv(c)
    except Exception:
        abort(400, 'could not create df')
    if df.isnull().values.any():
        abort(400, 'data contains nulls')
    cv_results = {}
    kf = KFold(n_splits=n_folds)
    x = df.drop(target, axis=1)
    y = df[target]
    best_mse = -1
    for coef in coefs:
        scores = []
        for train_index, test_index in kf.split(x):
            x_train, x_test = x.loc[train_index], x.loc[test_index]
            y_train, y_test = y.loc[train_index], y.loc[test_index]
            ridge = Ridge(alpha=coef, fit_intercept=fit_intercept)
            ridge.fit(x_train, y_train)
            y_pred = ridge.predict(x_train)
            mse_train_score = mse_score(y_train, y_pred)
            scores.append(mse_train_score)
            if best_mse == -1 or best_mse > mse_train_score:
                best_mse = mse_train_score
                best_coef = ridge.coef_
                best_intercept = ridge.intercept_
        cv_results[coef] = np.mean(scores)
        columns = x_train.columns
        weights = dict(zip(columns, best_coef))
        js = {'intercept': best_intercept, 'coef': weights}
        js = json.dumps(js)
        cv = []
        for key in cv_results:
            cv.append({'param_value': key, 'mean_mse': cv_results[key]})
        cv = json.dumps(cv)
        return js, cv


# prediction function
def predict(data, model):
    c = StringIO(data)
    try:
        df = pd.read_csv(c)
    except Exception:
        abort(400, 'could not create df')
    if df.isnull().values.any():
        abort(400, 'data contains nulls')
    inter = model['intercept']
    coefs = model['coef']
    heads = df.columns
    results = []
    for i in df.itertuples():
        p = 0
        for j in range(1, len(i)):
            p += i[j] * coefs[heads[j - 1]]
        results.append(p + inter)
    return results


def read_config():
    try:
        with open('config.json') as config_file:
            try:
                data = json.load(config_file)
                dsn = data["DB_DSN"]
                return dsn
            except ValueError:
                logging.exception("could not read info from config")
                sys.exit(1)
    except Exception:
        logging.exception("config file not found")
        sys.exit(1)


try:
    engine = create_engine(read_config())
    session_factory = sessionmaker(bind=engine)

    DeclarativeBase = declarative_base()
except Exception:
    logging.exception("failed engine creation")
    sys.exit(1)


class Models(DeclarativeBase):
    __tablename__ = 'models'

    id = Column(Integer, primary_key=True, autoincrement=True)
    model = Column(JSONB)
    cv_results = Column(JSONB)


def makedb():
    DeclarativeBase.metadata.create_all(engine)


app = Flask(__name__)

try:
    makedb()
except Exception:
    logging.exception("could not create db")
    sys.exit(1)


@app.errorhandler(Exception)
def error_handler(exc):
    if isinstance(exc, HTTPException):
        return jsonify({'error': str(exc.description)}), exc.code
    return jsonify({'error': "internal error"}), 500


@app.route('/train', methods=['POST'])
def train():
    d = request.get_json()
    try:
        data = str(d['data'])
        target = str(d['target'])
        n_folds = int(d['n_folds'])
        fit_intercept = bool(['fit_intercept'])
        l2_coef = list(d['l2_coef'])
    except Exception as exc:
        abort(400, f"invalid data: {exc}")
        return
    js, cv = model_trainig(data, target, n_folds, fit_intercept, l2_coef)
    session = session_factory()
    logging.info("Model training started")
    model = Models(model=js, cv_results=cv)
    logging.info("Model training completed")
    session.add(model)
    session.commit()
    logging.info("Model added to db")
    return jsonify({"model_id": model.id})


@app.route('/model/<int:model_id>', methods=['GET'])
def model_get(model_id):
    session = session_factory()
    model = session.query(Models).filter_by(id=model_id).first()
    if model is None:
        abort(404, f"not found: {model_id}")
        return
    logging.info("Model found")
    return jsonify({"model": model.model, "cv_results": model.cv_results})


@app.route('/predict', methods=['POST'])
def model_predict():
    d = request.get_json()
    try:
        data = str(d['data'])
        model_id = int(d['model_id'])
    except Exception as exc:
        abort(400, f"invalid data: {exc}")
        return
    session = session_factory()
    model_info = session.query(Models).filter_by(id=model_id).first()
    if model_info is None:
        abort(404, f"not found: {model_id}")
        return
    logging.info("Model found. Starting prediction")
    model = eval(model_info.model)
    res = predict(data, model)
    logging.info("Prediction done")
    return jsonify({"result": res})


if __name__ == '__main__':
    app.run(debug=False)
