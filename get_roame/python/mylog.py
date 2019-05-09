#!/usr/bin/env python
# -*- coding:utf-8 -*-

import logging
import getpass


class MyLog(object):
    def __init__(self):
        self.user = getpass.getuser()
        self.logger = logging.getLogger(self.user)
        self.logger.setLevel(logging.DEBUG)

        self.logfile = 'mylog.log'
        self.formatter = logging.Formatter('%(asctime)s %(levelname)s \
        %(name)s %(message)s \n')

        self.log_hand = logging.FileHandler(self.logfile, encoding='utf8')
        self.log_hand.setLevel(logging.DEBUG)
        self.log_hand.setFormatter(self.formatter)

        self.log_hand_st = logging.StreamHandler()
        self.log_hand_st.setLevel(logging.DEBUG)
        self.log_hand_st.setFormatter(self.formatter)

        self.logger.addHandler(self.log_hand)
        self.logger.addHandler(self.log_hand_st)


    def debug(self, msg):
        self.logger.debug(msg)


    def info(self, msg):
        self.logger.info(msg)


    def warn(self, msg):
        self.logger.warn(msg)


    def error(self, msg):
        self.logger.error(msg)


    def critical(self, msg):
        self.logger.critical(msg)


if __name__ == '__main__':
    mylog = MyLog()

    mylog.debug(u'test debug 呵')
    mylog.info(u'test info 呵')
    mylog.warn(u'test warn 呵')
    mylog.error(u'test error 呵')
    mylog.critical(u'test critical 呵')
