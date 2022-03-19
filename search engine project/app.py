from flask import Flask, request, render_template, url_for
import forsearch  # another .py file
import pandas as pd
from gensim.models import KeyedVectors
import pickle

# import all necessary files
app = Flask(__name__)
tfidf_corpus = forsearch.get_pickle('tfidf.pkl')
bm25_mx = pd.read_csv('bm25.csv')
w2v = forsearch.get_pickle('w2vclass.pkl')
w2v_exp = forsearch.get_pickle('w2vexp.pkl')
answers = forsearch.get_pickle('answers.pkl')
answers = answers.to_numpy()
vectorizer = pickle.load(open('vectorizer.pkl', 'rb'))
model_file = 'araneum_none_fasttextcbow_300_5_2018.model'
model = KeyedVectors.load(model_file)


@app.route('/')
def index():
    urls = {'назад': url_for('index')}
    if request.args:
        query = request.args['query']
        name = request.args["choices-single-defaul"]
        links = search(query, name)
        return render_template('result_page.html', result_s=links, model=name, query=query, urls=urls)
    return render_template('index.html')


def search(query, name):
    search_method = name
    if search_method == 'TF-IDF':
        result = forsearch.get_result(forsearch.q_to_vec_tfidf(query, vectorizer), tfidf_corpus, answers)
    elif search_method == 'BM25':
        result = forsearch.get_result(forsearch.q_to_vec(query, bm25_mx), bm25_mx, answers)
    elif search_method == 'word2vec_mean_vec':
        result = forsearch.get_result(forsearch.get_vector_w2v(query,  model), w2v, answers)
    elif search_method == 'word2vec_matrix':
        result = forsearch.get_result_w2vexp(forsearch.create_q_matrix(query, model), w2v_exp, answers)
    else:
        result = forsearch.get_result(forsearch.q_to_vec(query, tfidf_corpus), tfidf_corpus, answers)
    return result


if __name__ == '__main__':
    app.run(debug=False)
