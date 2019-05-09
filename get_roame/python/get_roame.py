#-*- coding:utf-8 -*-

import requests
import os
from bs4 import BeautifulSoup
from mylog import *

USER_AGENTS = [
    "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
    "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
    "Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
    "Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
    "Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
    "Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
    "Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
    "Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
    "Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
    "Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
    "Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
    "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
    "Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
]




class GetPicture():
    def __init__(self):
		self.log = MyLog()

		self.exist_images = []
		for image_name in os.listdir('D:\\Image Spider\\roame\\'):
			name, _ = os.path.splitext(image_name)
			self.exist_images.append(name)



		self.headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/18.17763'}
		self.cookies = {
			'Hm_lvt_633fe378147652d6cb58809821524bec': '1552293894,1552294302,1552377763,1552462512',
			'uid': '255971',
			'upw': 'a0663602d3f2ce2f0194a855af395d19',
			'cmd': '2L6hwBaU%40FuaUNQ5G4IxoLwxoHNuLK5g',
			'Hm_lpvt_633fe378147652d6cb58809821524bec': '1552462543'
		}
		self.base_url = 'https://www.roame.net'
		name = input('Input Name: ')

		self.page_url = 'https://www.roame.net/index/' + name + '/images/index.html'
		self.log.info('########################################Page %s Start###########################################' % str(0))
		self.log.info('Get page: %s' % self.page_url)
		urls, names = self.get_image_url()
		self.get_images(urls, names)
		self.log.info('########################################Page %s End###########################################' % str(0))

		for i in range(1, 50):
			self.page_url = 'https://www.roame.net/index/' + name + '/images/index_' \
			+ str(i) + '.html'
			self.log.info('########################################Page %s Start###########################################' % str(i))
			self.log.info('Get page: %s' % self.page_url)

			urls, names = self.get_image_url()
			self.get_images(urls, names)
			self.log.info('########################################Page %s End###########################################' % str(i))


    def get_images(self, urls, names):
        for name, url in zip(names, urls):
            if name in self.exist_images:
                self.log.warn('Image %s has existed, pass!' % name)
                continue
            self.log.info('Get image name: %s, url: %s' % (name, url))
            image = self.get_response_content(url)
            self.save_image(name, image)



    def get_image_url(self):
		html_content = self.get_response_content(self.page_url)
		soup = BeautifulSoup(html_content, 'lxml')

		img_urls = []
		img_names = []

		tmps = soup.find_all('div', attrs={'class': 'fbi'})
		for tmp in tmps:
			name = tmp.a.img['src']
			# https://ios.roame.net/files/J13fZpvLqGemQFxeFJbCgPP1ZU4gnf25rquZoYiq2Db/
			# ROAME_310416_EEC30F66.256.jpg
			tmp_list = name.split('/')
			name = tmp_list[-1].split('.')
			name = name[0]
			img_names.append(name)

			url = self.base_url + tmp.a['href']
			html_content = self.get_response_content(url)
			soup = BeautifulSoup(html_content, 'lxml')
			url = soup.find('div', attrs={'id': 'darlnks'})
			url = url.find_all('a')[1]
			url = url['href']
			url = url + name + '.jpg'
			print url
			img_urls.append(url)

			self.log.info('Add image url %s' % name)
			self.log.info('url %s' % url)


		return img_urls, img_names


    def get_response_content(self, url):
        try:
            response = requests.get(url=url, headers=self.headers, cookies=self.cookies)
        except:
            self.log.warn('url %s open error!' % url)
            return None
        else:
            self.log.info('url %s open success!' % url)
        return response.content


    def save_text(self, text_name, text):
        with open(text_name, 'w') as fp:
            fp.write(text)


    def save_image(self, image_name, image):
        try:
            with open('D:\\roame\\Image Spider\\' + image_name + '.jpg', 'wb') as fp:
                fp.write(image)
                self.log.info('Save image: %s successfully!' % image_name)
        except Exception as err:
            print 'save image error!'
            print err
            self.log.error('save image: %s error!' % image_name)


if __name__ == '__main__':
        GetPicture()