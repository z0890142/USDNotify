#!/bin/bash
DB_USER='developer';
DB_PASSWD='1qaz@WSX3edc';

mysql --user=$DB_USER --password=$DB_PASSWD << EOF
use ForeignCurrency

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,15_Heigest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 15 Day) and CURDATE() 
    and Price=fr.15_Heigest
) as f
on fr.SN=f.SN
set fr.15_Heigest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Month_Heigest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 3 Month) and CURDATE() 
    and Price=fr.3Month_Heigest
) as f
on fr.SN=f.SN
set fr.3Month_Heigest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,6Month_Heigest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 6 Month) and CURDATE() 
    and Price=fr.6Month_Heigest
) as f
on fr.SN=f.SN
set fr.6Month_Heigest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,1Year_Heigest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 1 Year) and CURDATE() 
    and Price=fr.1Year_Heigest
) as f
on fr.SN=f.SN
set fr.1Year_Heigest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Year_Heigest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.3Year_Heigest
) as f
on fr.SN=f.SN
set fr.3Year_Heigest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,5Year_Heigest FROM ForeignCurrencyBuyInPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 5 Year) and CURDATE() 
    and Price=fr.5Year_Heigest
) as f
on fr.SN=f.SN
set fr.5Year_Heigest_Date =f.Date;


update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,15_Lowest FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 15 Day) and CURDATE() 
    and Price=fr.15_Lowest
) as f
on fr.SN=f.SN
set fr.15_Lowest_Date =f.Date;


update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Month_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 3 Month) and CURDATE() 
    and Price=fr.3Month_Lowest
) as f
on fr.SN=f.SN
set fr.3Month_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,6Month_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 6 Month) and CURDATE() 
    and Price=fr.6Month_Lowest
) as f
on fr.SN=f.SN
set fr.6Month_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,1Year_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 1 Year) and CURDATE() 
    and Price=fr.1Year_Lowest
) as f
on fr.SN=f.SN
set fr.1Year_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,3Year_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 3 Year) and CURDATE() 
    and Price=fr.3Year_Lowest
) as f
on fr.SN=f.SN
set fr.3Year_Lowest_Date =f.Date;

update ForeignCurrencyRecord as fr inner join
(
    SELECT fsp.SN,Date,5Year_Lowest_Date FROM ForeignCurrencySellPrice as fsp
    join ForeignCurrencyRecord as fr on fsp.SN=fr.SN
    where fsp.Date between DATE_SUB(CURDATE(), INTERVAL 5 Year) and CURDATE() 
    and Price=fr.5Year_Lowest
) as f
on fr.SN=f.SN
set fr.5Year_Lowest_Date =f.Date;

EOF