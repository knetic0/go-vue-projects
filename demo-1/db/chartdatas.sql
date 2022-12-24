CREATE TABLE IF NOT EXISTS chartdatas(
	id SERIAL PRIMARY KEY,
	bookname VARCHAR(100) NOT NULL,
	booktype VARCHAR(100) NOT NULL,
	author VARCHAR(100) NOT NULL,
	popularity INT NOT NULL,
	totalbook INT NOT NULL
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Olasılıksız', 'Polisiye', 'Adam Fawer', 164000, 2907
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Suç ve Ceza', 'Roman', 'Fyodor Dostoyevski', 417000, 25400
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Açlık', 'Psikolojik', 'Knut Hamsun', 416000, 2100
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Saklambaç', 'Korku-Gerilim', 'N.G. Kabal', 158000, 865
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Aşk ve Gurur', 'Aşk', 'Jane Austen', 598000, 3339
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Savaş ve Barış', 'Tarih', 'Tolstoy', 231000, 11600
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Küçük İşler Büyük Özgürlükler', 'Kişisel Gelişim', 'Mert Başaran', 210000, 3292
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook) VALUES(
	'Oğuz Kağan Destanı', 'Öykü', 'Bilgin Adalı', 130000, 9500
);

SELECT * FROM chartdatas;