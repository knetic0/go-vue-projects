CREATE TABLE IF NOT EXISTS chartdatas(
	id SERIAL PRIMARY KEY,
	bookname VARCHAR(100) NOT NULL,
	booktype VARCHAR(100) NOT NULL,
	author VARCHAR(100) NOT NULL,
	popularity INT NOT NULL,
	totalbook INT NOT NULL,
	agelimit INT NULL
);

DROP TABLE chartdatas;

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Olasılıksız', 'Polisiye', 'Adam Fawer', 164000, 2907, 18
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Suç ve Ceza', 'Roman', 'Fyodor Dostoyevski', 417000, 25400, 18
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Açlık', 'Psikolojik', 'Knut Hamsun', 416000, 2100, 18
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook,agelimit) VALUES(
	'Saklambaç', 'Korku-Gerilim', 'N.G. Kabal', 158000, 865, 18
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Aşk ve Gurur', 'Aşk', 'Jane Austen', 598000, 3339, 13
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Savaş ve Barış', 'Tarih', 'Tolstoy', 231000, 11600, 13
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Küçük İşler Büyük Özgürlükler', 'Kişisel Gelişim', 'Mert Başaran', 210000, 3292, 13
);

INSERT INTO chartdatas(bookname, booktype, author, popularity, totalbook, agelimit) VALUES(
	'Oğuz Kağan Destanı', 'Öykü', 'Bilgin Adalı', 130000, 9500, 13
);

SELECT * FROM chartdatas WHERE agelimit >= 18;

SELECT * FROM chartdatas;