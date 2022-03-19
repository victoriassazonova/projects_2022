import mysql.connector

def create_db(host_address, port_address, username, user_password):
    """Creates db

    Parameters:
        host_address:str - host for db connection
        port_address:int - port for db connection
        username:str -  user for db connection
        user_password:str -  user password for db connection

    Returns:
        cursor
    """
    con = mysql.connector.connect(host=host_address, port=port_address, user=username, password=user_password)
    cur = con.cursor(dictionary=True)
    cur.execute("CREATE DATABASE IF NOT EXISTS chatbot;")
    con = mysql.connector.connect(host=host_address, port=port_address, database='chatbot',
                                  user=username, password=user_password)
    cur = con.cursor(dictionary=True)
    cur.execute("""CREATE TABLE IF NOT EXISTS chatbot_answers(
    id INT NOT NULL AUTO_INCREMENT,
    text_time DATETIME COMMENT 'Time of the message', 
    user_text VARCHAR(4) COMMENT 'Emoji of the user or TEXT if not an emoji',
    bot_answer VARCHAR(100) COMMENT 'Bot answer',
    PRIMARY KEY (id),
    UNIQUE INDEX id_UNIQUE (id ASC) VISIBLE)
    COMMENT 'Table for user texts and bot answers with time';""")
    return cur


def insert_into_db(text_time, user_text, bot_answer, host_address, port_address, username, user_password):
    """Writes to db

    Parameters:
        text_time:datetime -  time of the message
        user_text:str - user inputed text
        bot_answer:str -  bots answer
        host_address:str - host for db connection
        port_address:int - port for db connection
        username:str -  user for db connection
        user_password:str -  user password for db connection
    """
    con = mysql.connector.connect(host=host_address, port=port_address, database='chatbot', user=username,
                                  password=user_password)
    cur = con.cursor(dictionary=True)
    insert_line = f"INSERT INTO chatbot_answers (text_time, user_text, bot_answer) VALUES " \
                  f"('{text_time}', '{user_text}', '{bot_answer}')"
    cur.execute(insert_line)
    con.commit()
