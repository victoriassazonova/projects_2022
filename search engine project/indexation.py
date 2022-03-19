from sklearn.feature_extraction.text import CountVectorizer, TfidfVectorizer
import pickle
from math import log
import numpy as np
from sklearn.model_selection import train_test_split
import pandas as pd
from pymorphy2 import MorphAnalyzer
from nltk.corpus import stopwords
from gensim.models import KeyedVectors
import tarfile

stop_words = set(stopwords.words('russian'))


# preprocessing function
def preprocess(rawtext):
    ans = ''
    for char in rawtext:
        if char.isalpha() or char == ' ':
            ans += char
        else:
            ans += ' '
    words = ans.split(" ")
    w = []
    for i in words:
        if i != '':
            p = morph.parse(i.lower())[0].normal_form
            if p not in stop_words:
                w.append(p)
    return " ".join(w)


morph = MorphAnalyzer()


# function for corpora creation
def work_files():
    answers = pd.read_excel("answers_base.xlsx")
    questions = pd.read_excel("queries_base.xlsx")
    reply_base = pd.DataFrame([answers['Номер связки'], answers["Текст ответа"]]).transpose()
    reply_base = reply_base.to_numpy()
    d = {}
    for i in reply_base:
        d[i[0]] = i[1]
    train, test1 = train_test_split(questions, test_size=0.3)
    ans = pd.concat([pd.Series(row['Номер связки'], row["Текст вопросов"].split('\n'))
                     for _, row in answers.iterrows()]).reset_index()
    new = pd.DataFrame([train['Номер связки\n'], train["Текст вопроса"]]).transpose()
    new = new.rename(columns={'Номер связки\n': 'i', "Текст вопроса": 'k'}, inplace=False)
    ans = ans.rename(columns={0: 'i', "index": 'k'}, inplace=False)
    train = new.append(ans, ignore_index=True)
    train = train.dropna()
    train['f'] = train['i'].apply(lambda x: d[x])
    answers = pd.DataFrame([train['i'], train['f']]).transpose()
    train['proc'] = train['k'].apply(lambda x: preprocess(x))
    return answers, train['proc']


# functions for different matrix types
def make_tf_idf(corpus):
    vectorizer = TfidfVectorizer()
    td_matrix = pd.DataFrame(vectorizer.fit_transform(corpus).A, columns=vectorizer.get_feature_names())
    return td_matrix, vectorizer


def bm25(tfmtrx, avgdl, N, q, idx, ld):
    k = 2.0
    b = 0.75
    n = np.count_nonzero(tfmtrx[q])
    tf = tfmtrx[q][idx]
    idf = log((N-n+0.5)/(n+0.5))
    score = idf*((tf*(k+1))/(tf+k*(1-b+b*ld/avgdl)))
    return score


def make_bm25(corpus):
    vectorizer = CountVectorizer()
    Y = vectorizer.fit_transform(corpus).toarray()
    names = vectorizer.get_feature_names()
    tfmtrx = pd.DataFrame(data=Y, columns=names)
    N = len(corpus)
    lds = [len(corp) for corp in corpus]
    avgdl = sum(lds)/len(lds)
    bm25mtrx = pd.DataFrame(columns=names)
    for i, text in enumerate(corpus):
        current = []
        t = text.split()
        for w in names:
            if w in t:
                ld = len(t)
                current.append(bm25(tfmtrx, avgdl, N, w, i, ld))
            else:
                current.append(0)
        bm25mtrx = bm25mtrx.append(pd.Series(current, index=bm25mtrx.columns), ignore_index=True)
    return bm25mtrx


def normalize_vec(v):
    return v / np.sqrt(np.sum(v ** 2))


def create_matrix(corpus, model):
    matrix = []
    for i in corpus:
        lemmas = i.split(' ')
        lemmas_vectors = np.zeros((len(lemmas), model.vector_size))
        for idx, lemma in enumerate(lemmas):
            try:
                if lemma in model:
                    lemmas_vectors[idx] = model[lemma]
            except:
                continue
        if lemmas_vectors.shape[0] is not 0:
            vec = np.mean(lemmas_vectors, axis=0)
            matrix.append(normalize_vec(vec))
    return np.array(matrix)


def create_doc_matrix(docs, model):
    matrix = []
    for text in docs:
        lemmas = text.split(" ")
        lemmas_vectors = np.zeros((len(lemmas), model.vector_size))
        for idx, lemma in enumerate(lemmas):
            try:
                if lemma in model:
                    lemmas_vectors[idx] = normalize_vec(model[lemma])
            except:
                continue
        matrix.append(lemmas_vectors)
    return matrix


# making all the necessary files
def make_files():
    tar = tarfile.open("araneum_none_fasttextcbow_300_5_2018.tgz", "r")
    tar.extractall()
    model_file = 'araneum_none_fasttextcbow_300_5_2018.model'

    answers, train = work_files()

    model = KeyedVectors.load(model_file)

    tf_matrix, vectorizer = make_tf_idf(train)
    output = open('tfidf.pkl', 'wb')
    pickle.dump(tf_matrix, output)
    output.close()
    output = open('vectorizer.pkl', 'wb')
    pickle.dump(vectorizer, output)
    output.close()

    bm25_matrix = make_bm25(train)
    bm25_matrix.to_csv('bm25.csv', index=False)

    w2v_matrix = create_matrix(train, model)
    output = open('w2vclass.pkl', 'wb')
    pickle.dump(w2v_matrix, output)
    output.close()

    w2v_matrix = create_doc_matrix(train, model)
    output = open('w2vexp.pkl', 'wb')
    pickle.dump(w2v_matrix, output)
    output.close()

    output = open('answers.pkl', 'wb')
    pickle.dump(answers, output)
    output.close()


def main():
    make_files()


if __name__ == "__main__":
    main()
