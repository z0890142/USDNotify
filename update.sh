#!/bin/bash
DB_USER='developer';
DB_PASSWD='1qaz@WSX3edc';

mysql --user=$DB_USER --password=$DB_PASSWD << EOF

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,15_Heighest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.15_Heighest
) as f
on fr.SN=f.SN
set fr.15_Heighest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Month_Heighest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.3Month
) as f
on fr.SN=f.SN
set fr.3Month_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,6Month_Heighest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.6Month
) as f
on fr.SN=f.SN
set fr.6Month_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,1Year_Heighest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.1Year
) as f
on fr.SN=f.SN
set fr.1Year_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Year_Heighest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.3Year
) as f
on fr.SN=f.SN
set fr.3Year_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,5Year_Heighest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.5Year
) as f
on fr.SN=f.SN
set fr.5Year_Date =f.Date;


update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,15_Lowest FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.15_Lowest
) as f
on fr.SN=f.SN
set fr.15_Lowest_Date =f.Date;


update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Month_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.3Month_Lowest
) as f
on fr.SN=f.SN
set fr.3Month_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,6Month_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.6Month_Lowest
) as f
on fr.SN=f.SN
set fr.6Month_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,1Year_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.1Year_Lowest
) as f
on fr.SN=f.SN
set fr.1Year_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Year_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.3Year_Lowest
) as f
on fr.SN=f.SN
set fr.3Year_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,5Year_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord fr on fsp.SN=fr.SN
    where Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.5Year_Lowest
) as f
on fr.SN=f.SN
set fr.5Year_Lowest_Date =f.Date;

EOF