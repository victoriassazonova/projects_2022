import pickle
import indexation # another .py file
import numpy as np


# all functions for the search engine
def get_pickle(filename):
    pickle_in = open(filename, "rb")
    matrix = pickle.load(pickle_in)
    return matrix


def q_to_vec(text, matrix):
    ar = indexation.preprocess(text).split(' ')
    vector = [0] * len(list(matrix.columns))
    for i, word in enumerate(list(matrix.columns)):
        if word in ar:
            vector[i] = 1
    return vector


def q_to_vec_tfidf(text, vectorizer):
    ar = []
    ar.append(indexation.preprocess(text))
    vec = vectorizer.transform(ar).toarray()
    return vec[0]


def get_vector_w2v(query, model):
    lemmas = indexation.preprocess(query).split(" ")
    lemmas_vectors = np.zeros((len(lemmas), model.vector_size))
    vec = np.zeros((model.vector_size,))
    for idx, lemma in enumerate(lemmas):
        try:
            if lemma in model:
                lemmas_vectors[idx] = model[lemma]
        except:
            continue
    if lemmas_vectors.shape[0] is not 0:
        vec = np.mean(lemmas_vectors, axis=0)
    return indexation.normalize_vec(vec)


def create_q_matrix(text, model):
    lemmas = indexation.preprocess(text).split(" ")
    lemmas_vectors = np.zeros((len(lemmas), model.vector_size))
    for idx, lemma in enumerate(lemmas):
        try:
            if lemma in model:
                lemmas_vectors[idx] = indexation.normalize_vec(model[lemma])
        except:
            continue
    return lemmas_vectors


def get_result(vector, matrix, answers):
    result_array = np.dot(matrix, vector)
    sorted_result_array = sorted([(e, i) for i, e in enumerate(result_array)], reverse=True)
    search_result = []
    s = []
    for item in list(sorted_result_array)[:-15]:
        if answers[item[1]][0] not in s:
            s.append(answers[item[1]][0])
            search_result.append(answers[item[1]])
        if len(s) == 5:
            break
    return search_result


def search(docs, query, reduce_func=np.max, axis=0):
    sims = []
    for doc in docs:
        sim = doc.dot(query.T)
        sim = reduce_func(sim, axis=axis)
        sims.append(sim.sum())
    sort_index = np.argsort(sims)[::-1]
    return sort_index


def get_result_w2vexp(vector, matrix, answers):
    result_array = search(matrix, vector)
    search_result = []
    s=[]
    for item in result_array[:-15]:
        if answers[item][0] not in s:
            s.append(answers[item][0])
            search_result.append(answers[item])
        if len(s) == 5:
            break
    return search_result

