import chat_bot
import data_base

host = '127.0.0.1'
port = 3306
username = 'root'
password = '{Password}'


def main():
    data_base.create_db(host, port, username, password)
    chat_bot.chat_bot(host, port, username, password)


if __name__ == '__main__':
    main()
