DROP TABLE IF EXISTS "manga";
CREATE TABLE "manga" (
    "manga_id" SERIAL PRIMARY KEY NOT NULL,
    "title" VARCHAR(255) NOT NULL,
    "english_title" VARCHAR(255),
    "japanese_title" VARCHAR(255),
    "author" VARCHAR(100) NOT NULL,
    "artist" VARCHAR(100), 
    "status" smallint NOT NULL DEFAULT '0',
    "published_on" date NOT NULL,
    "finished_on" date DEFAULT NULL,
    "synopsis" text,
    "user_id" int not null
);

INSERT INTO "manga" VALUES 
(1, 'Berserk','Berserk','ベルセルク', 'Miura Kentaro', 'Miura Kentaro, Studio Gaga', 0, '1989-08-25', NULL,'Guts, a former mercenary now known as the "Black Swordsman," is out for revenge. After a tumultuous childhood, he finally finds someone he respects and believes he can trust, only to have everything fall apart when this person takes away everything important to Guts for the purpose of fulfilling his own desires. Now marked for death, Guts becomes condemned to a fate in which he is relentlessly pursued by demonic beings.
Setting out on a dreadful quest riddled with misfortune, Guts, armed with a massive sword and monstrous strength, will let nothing stop him, not even death itself, until he is finally able to take the head of the one who stripped him—and his loved one—of their humanity.',1
),
(2, 'Act-Age','Act-age','アクタージュ act-age','Matsuki, Tatsuya', 'Usazaki, Shiro', 3, '2018-01-22', '2020-08-11', 
'Ever since she was a child, Kei Yonagi has been attracted to the idea of acting, as she has always attentively watched how actions and emotions differ depending on the type of character portrayed. Now a teenager, she is broke, and is the only one able to financially support her young siblings. She starts auditioning for any role possible, without much success.
Her luck begins to change when Sumiji Kuroyama, a renowned director, sees Yonagi`s performance and is amazed by her ability to immerse herself in the role—the so-called method acting. Although she is a diamond in the rough, if she keeps using the same technique over and over, it could lead to her being severely damaged. So, he makes the decision to help polish her, giving Yonagi the opportunity she needs. However, the road ahead of her is far from easy, and she will need to adapt to overcome the various challenges that come her way.',1
),
(3, '5-toubun no Hanayome','The Quintessential Quintuplets','五等分の花嫁','Negi Haruba','Negi Haruba',1,'2017-08-09','2020-02-19','Considered a genius, high schooler Fuutarou Uesugi excels at studying and obtains a perfect score on every test. Due to his intense focus on that regard, he is a reclusive person with no friends. Additionally, he lives in a tight financial state as a result of family debts.
Fuutarou`s mundane lifestyle is interrupted when Itsuki Nakano, a new transfer student, contests him for his usual lunch seat. After a short altercation, Fuutarou emerges victorious by insulting Itsuki`s eating habits, which angers her enough to leave. However, when Fuutarou learns that he has been offered the private tutor position of a wealthy family`s academically hopeless daughters, he immediately regrets his prior encounter. It turns out that the beneficiaries of his tutoring are none other than Itsuki and her four identical siblings: the shy Miku, the cheerful Yotsuba, the devious Nino, and the mature Ichika!
Teaching these quintuplets may prove more difficult for Fuutarou than he initially expected, as the last thing they want to do is seriously study. Now, Fuutarou must face the various challenges of tutoring five beautiful yet eccentric girls.',1
);

DROP TABLE IF EXISTS "user";
CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY NOT NULL,
    "name" VARCHAR(20) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(20) NOT NULL,
    "token" VARCHAR(255),
    "manga" text
);

INSERT INTO "user" ("id", "name", "email","password", "manga") VALUES
(1, 'Karina', 'Karina@gmail.com', 'karina123','Berserk, Act-Age, 5-toubun no Hanayome');

CREATE INDEX "manga_FK" ON "user"("manga");
CREATE INDEX "email_INDEX" ON "user"("email"); 
ALTER TABLE "manga" ADD CONSTRAINT "user_FK" FOREIGN KEY ("user_id") REFERENCES "user" ("id");
