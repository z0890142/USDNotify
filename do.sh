#!/bin/bash
DB_USER='developer';
DB_PASSWD='1qaz@WSX3edc';

mysql --user=$DB_USER --password=$DB_PASSWD << EOF

use ForeignCurrency

update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Max(Price),3) as price FROM ForeignCurrencyBuyInPrice where Date between DATE_SUB(CURDATE(), INTERVAL 15 Day) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.15_Heigest =f.price;

update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Max(Price),3) as price FROM ForeignCurrencyBuyInPrice where Date between DATE_SUB(CURDATE(), INTERVAL 3 MONTH) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.3Month_Heigest =f.price;

update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Max(Price),3) as price FROM ForeignCurrencyBuyInPrice where Date between DATE_SUB(CURDATE(), INTERVAL 6 MONTH) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.6Month_Heigest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Max(Price),3) as price FROM ForeignCurrencyBuyInPrice where Date between DATE_SUB(CURDATE(), INTERVAL 1 Year) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.1Year_Heigest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Max(Price),3) as price FROM ForeignCurrencyBuyInPrice where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.3Year_Heigest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Max(Price),3) as price FROM ForeignCurrencyBuyInPrice where Date between DATE_SUB(CURDATE(), INTERVAL 5 Year) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.5Year_Heigest =f.price;


update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Min(Price),3) as price FROM ForeignCurrencySellPrice where Date between DATE_SUB(CURDATE(), INTERVAL 15 Day) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.15_Lowest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Min(Price),3) as price FROM ForeignCurrencySellPrice where Date between DATE_SUB(CURDATE(), INTERVAL 3 MONTH) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.3Month_Lowest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Min(Price),3) as price FROM ForeignCurrencySellPrice where Date between DATE_SUB(CURDATE(), INTERVAL 6 MONTH) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.6Month_Lowest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Min(Price),3) as price FROM ForeignCurrencySellPrice where Date between DATE_SUB(CURDATE(), INTERVAL 1 Year) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.1Year_Lowest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Min(Price),3) as price FROM ForeignCurrencySellPrice where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.3Year_Lowest =f.price;
update ForeignCurrencyRecord as fr inner join
(SELECT SN,round(Min(Price),3) as price FROM ForeignCurrencySellPrice where Date between DATE_SUB(CURDATE(), INTERVAL 5 Year) and CURDATE() group by SN) as f
on fr.SN=f.SN
set fr.5Year_Lowest =f.price;

EOF

