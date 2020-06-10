import matplotlib
import matplotlib.pyplot as plt
import datetime
import time
from dateutil.relativedelta import relativedelta
import mysql.connector
from mysql.connector import Error
import seaborn as sns
import pandas as pd
import matplotlib.ticker as ticker
import sys
import configparser


matplotlib.use('Agg')


price=[]
date=[]
name=sys.argv[1]
SN=sys.argv[2]

try:
    config = configparser.ConfigParser()
    
    #relative path of main.go or ABSOLUTE_PATH
    config.read('./config/config.ini') 

    connection = mysql.connector.connect(
        host=config.get('default', 'host'),          # 主機名稱
        database=config.get('default', 'database'), # 資料庫名稱
        user=config.get('default', 'user'),        # 帳號
        password=config.get('default', 'pwd'),
        auth_plugin='mysql_native_password'
    )
    
    # 查詢資料庫
    cursor = connection.cursor()
    cursor.execute("select Price,Date from ForeignCurrencySellPrice where SN="+SN+" and Date between DATE_SUB(CURDATE(), INTERVAL 1 Year) and CURDATE()")

    # 取回全部的資料
    records = cursor.fetchall()
    print("資料筆數：", cursor.rowcount)

    # 列出查詢的資料
    for (Price,Date) in records:
        price.append(Price)
        date.append(str(Date)[0:10])
    df = pd.DataFrame({'price':price, 'date':date})
    df.head()
except Error as e:
    print("資料庫連接失敗：", e)

finally:
    if (connection.is_connected()):
        cursor.close()
        connection.close()
        print("資料庫連線已關閉")

end = datetime.date.today() #開始時間結束時間,選取最近一年的資料
start =  end-relativedelta(years=1)
end = end.strftime("%Y%m%d") 
start = start.strftime("%Y%m%d") 

sns.set_style("whitegrid")

f, ax = plt.subplots(figsize = (8, 6))
ax.set_title(name+" "+start+"-"+end, fontsize=18, position=(0.5,1.05))
ax.tick_params(axis='y',labelsize=8) # y轴
ax.tick_params(axis='x',labelsize=8, rotation=45) # x轴

ax.xaxis.set_major_locator(ticker.MultipleLocator(base=10))

sns.lineplot(x='date', y='price', data=df)

ax.set(xlabel='', ylabel='')

#relative path of main.go or ABSOLUTE_PATH
f.savefig("./static/picture/"+name+'.jpg', dpi=200, bbox_inches='tight')
