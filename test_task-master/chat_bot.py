import datetime
from utils import emotions, response
from data_base import insert_into_db


def chat_bot(host_address, port_address, username, user_password):
    """Cycle - user input/bots answer
    Writes everything to db

    Parameters:
        host_address:str - host for db connection
        port_address:int - port for db connection
        username:str -  user for db connection
        user_password:str -  user password for db connection
    """
    last_emotion = None
    while True:
        last_emotion, bot_answer, text_time, user_text = get_answer(last_emotion)
        print(bot_answer)
        insert_into_db(text_time, user_text, bot_answer, host_address, port_address, username, user_password)


def get_answer(last_emotion):
    """Gets bots answer based on users last two messages

    Parameters:
        last emotion:str - previous text

    Returns:
        str:emotion
        str: bots answer
        datetime: time of the message
        text: inputed text
    """
    current_emotion, text_time, user_text = get_current()
    if last_emotion is None and current_emotion is not None:
        response_emotion = '{} greeting'.format(current_emotion)
    elif last_emotion is not None and current_emotion is not None:
        response_emotion = '{} {}'.format(last_emotion, current_emotion)
    else:
        response_emotion = 'not understood'
    return current_emotion, response[response_emotion], text_time, user_text


def get_current():
    """Gets users input and gets its emotions
    based on the emotions dict from utils.py

    Returns:
        str:emotion (returns None if user inputed not an emoji)
        datetime: time of the message
        text: inputed text
    """
    raw_input = input()
    text_time = datetime.datetime.now()
    user_text = raw_input
    emoji = raw_input.encode('utf-8')
    if emoji in emotions.keys():
        return emotions[emoji], text_time, user_text
    else:
        # if user submitted a wrong input - not an emoji,
        # we don't keep his message as it might be too long
        return None, text_time, 'TEXT'
