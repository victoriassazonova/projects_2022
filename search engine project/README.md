## This is my 'covid-19'db search engine project with 4 different search methods

It is trained on a data set of covid-related questions and answers and allows to search for answers to your question.
Its made with flask, the front is Html, csv and java script.

The search methods:
* tf-idf
* okapi bm25
* word2vec - mean of all vectors
* word2vec - matrix for every text in the collection

To make this thing work you will need:
<ol>
<li>to download araneum_none_fasttextcbow_300_5_2018 model from RusVectores</li>
<li>to download the matrix files or the xlsx files, if you want to make the matrixes too https://drive.google.com/drive/folders/1Hy68J59zkAI7xjgw0nxw0QeNqJU1H_FD?usp=sharing</li>
<li>start the flask app</li>
</ol>