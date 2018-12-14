package handler

import (
	"github.com/gin-gonic/gin"
)

type t_activity struct {
	Activity_id       string
	Activity_type     int
	Dealer_id         string
	City_id           int
	Title             string
	Title_short       string
	Picture           string
	Start_date        string
	End_date          string
	Phone             string
	Activity_discount int
	Activity_address  string
	Detail            string
	Packages          string
	Car_model         string
	Activity_status   string
	Createtime        string
	Updatetime        string
}

func ActivityList(c *gin.Context) {
	result1 := t_activity{}
	result1.Activity_id = "4028832b4ec02287014ec82cb88b0021"
	result1.Activity_type = 2
	result1.Dealer_id = "S2015037182764"
	result1.City_id = 440100
	result1.Title = "广利楼兰预售会暨东南亚风情水果节"
	result1.Title_short = "订车获遥控飞机"
	result1.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150723/1437633148063.jpg"
	result1.Start_date = "2015-07-25 00:00:00"
	result1.End_date = "2015-07-31 15:26:21"
	result1.Phone = "020-86388988"
	result1.Activity_discount = 0
	result1.Activity_address = "广州市白云区下塘西路614号（雕塑公园正对面）"
	result1.Detail = `！！特大好消息！！

！！！所有人都要冷静！！！

本周六7月25日

东风日产广利白金店举办楼兰预售会暨东南亚风情水果节

届时还有汽车之家同台合作，强强联手，

周末拒绝无聊，一起来嗨！！！聚惠就是现在！


1、楼兰预售会暨东南亚风情水果节

①订车获赠招商银行1000元油卡（仅限活动当天前10名预订客户）

②置换新楼兰送1000元积分

3、保有客户专场

①缤纷雪糕免费品尝（可爱多、五羊、梦龙、阿波罗等）

②转介绍：保有客户送300代金劵 200积分

被推荐客户送500元礼品劵

合伙订车新老客户，各赠送水果篮一份；

③东风日产保有客户增购、置换东风NISSAN全系车型（`
	result1.Packages = `["%E5%88%B0%E5%BA%97%E6%91%87%E4%B8%80%E6%91%87%E8%B5%A2%E5%8F%96%E4%B8%B0%E5%AF%8C%E5%88%B0%E5%BA%97%E7%A4%BC","%E9%81%A5%E6%8E%A7%E9%A3%9E%E6%9C%BA","%E8%AE%A2%E8%BD%A6%E7%A0%B8%E9%87%91%E8%9B%8B%E5%BE%97%E4%B8%9C%E9%A3%8E%E6%97%A5%E4%BA%A7%E9%99%90%E9%87%8F%E7%89%88%E7%A4%BC%E5%93%81","%E9%99%90%E5%89%8D10%E5%90%8D%E8%AE%A2%E8%BD%A6%E5%AE%A2%E6%88%B7","%E5%87%A1%E8%AE%A2%E8%BD%A6%EF%BC%8C%E5%8D%B3%E5%8F%AF%E5%8F%82%E4%B8%8E5%E8%BD%AE%E5%BE%AA%E7%8E%AF%E6%83%8A%E5%96%9C%E5%A4%A7%E6%8A%BD%E5%A5%96%EF%BC%8C%E5%A4%9A%E9%87%8D%E4%B8%AD%E5%A5%96%E6%9C%BA%E4%BC%9A%EF%BC%9B","%E4%BA%AB%E5%8F%978000%E5%85%83%E8%A1%A5%E8%B4%B4","%E4%BA%AB%E5%8F%974000%E5%85%83%E8%A1%A5%E8%B4%B4","%E4%BA%AB%E5%8F%972000%E5%85%83%E8%A1%A5%E8%B4%B4"]`
	result1.Car_model = `["10060", "10220", "24526", "353861"]`
	result1.Activity_status = "1"
	result1.Createtime = "2015-09-19 11:43:11"
	result1.Updatetime = "2015-09-19 11:43:12"

	result2 := t_activity{}
	result2.Activity_id = "4028832b4ec94b0f014ece1595060006"
	result2.Activity_type = 1
	result2.Dealer_id = "S2015012829704"
	result2.City_id = 440100
	result2.Title = "2015新楼兰火热预售"
	result2.Title_short = "预订新楼兰 订金赚红利"
	result2.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437977286597.jpg"
	result2.Start_date = "2015-07-29 14:06:22"
	result2.End_date = "2015-08-08 18:00:00"
	result2.Phone = "020-85553008"
	result2.Activity_discount = 0
	result2.Activity_address = "中山大道西895号棠东（骏景花园对面）"
	result2.Detail = `1、预售期间订车，订金享8月8日新车上市当天余额宝的利息红利！
2、订车越早，定金越高，红利越多！
3、所得红利可用于选购新车精品活兑换等值市场礼品！
4、新车上市后，若不满意新车可无条件退订！红利同样可换购礼品！`
	result2.Packages = `["%E6%96%B9%E5%90%91%E7%9B%98%E5%A5%97","%E4%B9%B0%E8%BD%A6%E9%80%81%E8%BD%A6%E8%86%9C"]`
	result2.Car_model = `["353861"]`
	result2.Activity_status = "1"
	result2.Createtime = "2015-09-19 11:43:13"
	result2.Updatetime = "2015-09-19 11:43:14"

	result3 := t_activity{}
	result3.Activity_id = "4028832b4ed50893014ed8711f3c0004"
	result3.Activity_type = 1
	result3.Dealer_id = "S2015036970924"
	result3.City_id = 441300
	result3.Title = "奇骏最高优惠1.2万 店内现车充足"
	result3.Title_short = "奇骏直降1.2万"
	result3.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150901/1441089585326.jpg"
	result3.Start_date = "2015-07-29 00:00:00"
	result3.End_date = "2015-08-09 00:00:00"
	result3.Phone = "0752-2285272"
	result3.Activity_discount = 0
	result3.Activity_address = "广东省惠州市三环路与口岸路交接处（实验中学旁）"
	result3.Detail = `本店内奇骏现车销售，颜色可选，目前购车部分加装车型可优惠1.20万元，感兴趣的朋友可以到店咨询购买，咨询电话：0752-2285272`
	result3.Packages = `["%E6%96%B0%E5%A5%87%E9%AA%8F%E7%9B%B4%E9%99%8D1.2%E4%B8%87"]`
	result3.Car_model = `["47710"]`
	result3.Activity_status = "1"
	result3.Createtime = "2015-09-19 11:43:15"
	result3.Updatetime = "2015-09-19 11:43:16"

	result4 := t_activity{}
	result4.Activity_id = "4028832b4ed50893014ed87aa5890006"
	result4.Activity_type = 1
	result4.Dealer_id = "S2015036970924"
	result4.City_id = 441300
	result4.Title = "天籁现车充足 最高现金优惠1.9万元"
	result4.Title_short = "天籁直降1.9万"
	result4.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150808/1439036597808.jpg"
	result4.Start_date = "2015-07-30 00:00:00"
	result4.End_date = "2015-08-30 00:00:00"
	result4.Phone = "0752-2285272"
	result4.Activity_discount = 0
	result4.Activity_address = "广东省惠州市三环路与口岸路交接处（实验中学旁）"
	result4.Detail = `店内天籁现车销售，颜色可选，目前购车部分车型可优惠3.30万元，感兴趣的朋友可以到店咨询购买，热线4008300451`
	result4.Packages = `["%E5%A4%A9%E7%B1%81%E6%9C%80%E9%AB%98%E7%8E%B0%E9%87%91%E4%BC%98%E6%83%A01.9%E4%B8%87%E5%85%83"]`
	result4.Car_model = `["10220", "353861", "47710", "49291", "492911"]`
	result4.Activity_status = "1"
	result4.Createtime = "2015-09-19 11:43:17"
	result4.Updatetime = "2015-09-19 11:43:18"

	result5 := t_activity{}
	result5.Activity_id = "4028832b4ef205c5014ef27cd4190000"
	result5.Activity_type = 1
	result5.Dealer_id = "S2015012829704"
	result5.City_id = 440100
	result5.Title = "盛世经典 东风日产“千人限时特卖会”"
	result5.Title_short = "日产“千人限时特卖会”"
	result5.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437976084376.jpg"
	result5.Start_date = "2015-08-09 09:00:00"
	result5.End_date = "2015-08-09 18:00:00"
	result5.Phone = "020-85553008"
	result5.Activity_discount = 0
	result5.Activity_address = "广州市天河区中山大道西895号"
	result5.Detail = `第一重礼：预交诚意金可变3000元!
第二重礼：签到即送价值199大礼包!
第三重礼：新楼兰广州上市，前10名现场订楼兰，享全国优先提车!
第四重礼：现金一把抓，手有多大钱就有多少!
第五重礼：玩转大富翁，大奖任你赢!
第六重礼：精品一元拍，价格者得!
第七重礼：厂家赞助礼—终身延保，排队领取，先到先得，数量有限，送完即止!
第八重礼：总经理特别礼!
第九重礼：前30名订车客户额外获赠价值280元行车记录仪一台!
第十重礼：15周年庆典十轮狂欢礼，十轮大奖疯狂抽取，越早订车中奖越多!`
	result5.Packages = `["%E7%8E%B0%E5%9C%A8%E6%8A%A5%E5%90%8D%E5%8F%82%E5%8A%A0%E6%B4%BB%E5%8A%A8%E6%8A%A2%E5%85%88%E8%B5%A2%E5%8F%961000%E5%85%83%E6%B2%B9%E5%8D%A1%21"]`
	result5.Car_model = `["10060", "10140", "10220", "14660", "24526", "30606", "47710", "49190", "49291", "492911"]`
	result5.Activity_status = "1"
	result5.Createtime = "2015-09-19 11:43:19"
	result5.Updatetime = "2015-09-19 11:43:20"

	result6 := t_activity{}
	result6.Activity_id = "4028832b4ef92959014efd1596880004"
	result6.Activity_type = 1
	result6.Dealer_id = "S2015072210924"
	result6.City_id = 370700
	result6.Title = "青山启辰巅峰钜惠 八月送清凉"
	result6.Title_short = "青山启辰巅峰钜惠"
	result6.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437967984964.jpg"
	result6.Start_date = "2015-08-08 15:00:00"
	result6.End_date = "2015-08-09 23:55:00"
	result6.Phone = "0536-8369002"
	result6.Activity_discount = 0
	result6.Activity_address = "山东省潍坊市潍城区殷大路与北宫西街交叉口300米（潍城区政府以西3公里处）"
	result6.Detail = `青江日产8月8日巅峰购车第四季闪亮登场！东风日产的粉丝们！这是你们享受低价的购车好机会！假如您在我们之前的活动中没考虑好，又或者是在之前的优惠中犹豫了，再或者是一不小心错过了，想买但是又找不到合适的机会，没关系，东风日产潍坊青江专营店8月8日钜惠来袭！狂购盛宴、巅峰钜惠，满足您便捷、低价购车的需求！您无需周旋于各大4S店，无需费尽力气讨价还价，无需提出N种买送要求，8月8日青江日产，您想要的，我们全都给你！`
	result6.Packages = `["%E8%BF%9B%E5%BA%97%E7%AD%BE%E5%88%B0%E5%B0%B1%E5%8F%AF%E5%8F%82%E4%B8%8E%E5%A4%A7%E5%AF%8C%E7%BF%81%E6%B8%B8%E6%88%8F%EF%BC%8C%E7%96%AF%E7%8B%82%E9%80%81%E8%B1%AA%E7%A4%BC%E5%85%88%E5%88%B0%E5%85%88%E5%BE%97%EF%BC%9B","%E7%8E%B0%E5%9C%BA%E8%AE%A2%E8%BD%A6%E5%B0%B1%E9%80%81%E5%8D%83%E5%85%83%E8%BD%A6%E8%A3%85%E5%A4%A7%E7%A4%BC%E5%8C%85%EF%BC%9B","%E7%8E%B0%E5%9C%BA%E8%AE%A2%E8%BD%A6%E5%AE%A2%E6%88%B7%E5%9D%87%E5%8F%AF%E5%8F%82%E4%B8%8E%E7%A0%B8%E9%87%91%E8%9B%8B%E8%B5%A2%E5%A5%BD%E7%A4%BC%EF%BC%8C%E7%B2%BE%E7%BE%8E%E7%A4%BC%E5%93%81%E9%9A%8F%E5%BF%83%E6%8A%BD","%E8%B4%AD%E8%BD%A6%E5%8D%B3%E6%9C%89%E6%9C%BA%E4%BC%9A%E8%8E%B7%E5%BE%971800%E5%85%83%E7%BE%8E%E7%9A%84%E9%AB%98%E7%AB%AF%E5%AE%B6%E7%94%B5%E5%9B%9B%E4%BB%B6%E5%A5%97%EF%BC%9B","%E7%8E%B0%E5%9C%BA%E8%AE%A2%E8%BD%A6%E5%AE%A2%E6%88%B7%E9%99%A4%E4%BB%A5%E4%B8%8A%E7%A4%BC%E5%93%81%E5%A4%96%EF%BC%8C%E8%BF%98%E5%8F%AF%E8%8E%B7%E5%BE%97%E7%BA%AA%E5%BF%B5%E7%A4%BC%E5%93%81%E4%B8%80%E4%BB%BD%EF%BC%9B","%E8%80%81%E5%AE%A2%E6%88%B7%E6%8E%A8%E8%8D%90%E6%96%B0%E5%AE%A2%E6%88%B7%E6%88%90%E5%8A%9F%E8%B4%AD%E8%BD%A6%EF%BC%8C%E5%8D%B3%E9%80%81%E7%9B%B8%E5%BA%94%E7%A7%AF%E5%88%86%EF%BC%8C%E7%A7%AF%E5%88%86%E5%BD%93%E9%92%B1","%E8%BF%9B%E5%BA%97%E8%AF%84%E4%BC%B0%E4%BA%8C%E6%89%8B%E8%BD%A6%E3%80%81%E5%9B%9E%E5%8E%82%E4%BF%9D%E5%85%BB%E7%88%B1%E8%BD%A6%EF%BC%8C%E7%B2%BE%E7%BE%8E%E7%A4%BC%E5%93%81%E9%9A%8F%E5%BF%83%E4%BA%AB%EF%BC%81"]`
	result6.Car_model = `["43013", "430131", "430132", "49299", "49300"]`
	result6.Activity_status = "1"
	result6.Createtime = "2015-09-19 11:43:21"
	result6.Updatetime = "2015-09-19 11:43:22"

	results := []t_activity{result1, result2, result3, result4, result5, result6}
	c.JSON(200, gin.H{"result": results})
}

func ActivityById(c *gin.Context) {
	id := c.Params.ByName("id")
	result := t_activity{}

	if id == "4028832b4ec02287014ec82cb88b0021" {
		result1 := t_activity{}
		result1.Activity_id = "4028832b4ec02287014ec82cb88b0021"
		result1.Activity_type = 2
		result1.Dealer_id = "S2015037182764"
		result1.City_id = 440100
		result1.Title = "广利楼兰预售会暨东南亚风情水果节"
		result1.Title_short = "订车获遥控飞机"
		result1.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150723/1437633148063.jpg"
		result1.Start_date = "2015-07-25 00:00:00"
		result1.End_date = "2015-07-31 15:26:21"
		result1.Phone = "020-86388988"
		result1.Activity_discount = 0
		result1.Activity_address = "广州市白云区下塘西路614号（雕塑公园正对面）"
		result1.Detail = `！！特大好消息！！

！！！所有人都要冷静！！！

本周六7月25日

东风日产广利白金店举办楼兰预售会暨东南亚风情水果节

届时还有汽车之家同台合作，强强联手，

周末拒绝无聊，一起来嗨！！！聚惠就是现在！


1、楼兰预售会暨东南亚风情水果节

①订车获赠招商银行1000元油卡（仅限活动当天前10名预订客户）

②置换新楼兰送1000元积分

3、保有客户专场

①缤纷雪糕免费品尝（可爱多、五羊、梦龙、阿波罗等）

②转介绍：保有客户送300代金劵 200积分

被推荐客户送500元礼品劵

合伙订车新老客户，各赠送水果篮一份；

③东风日产保有客户增购、置换东风NISSAN全系车型（`
		result1.Packages = `["%E5%88%B0%E5%BA%97%E6%91%87%E4%B8%80%E6%91%87%E8%B5%A2%E5%8F%96%E4%B8%B0%E5%AF%8C%E5%88%B0%E5%BA%97%E7%A4%BC","%E9%81%A5%E6%8E%A7%E9%A3%9E%E6%9C%BA","%E8%AE%A2%E8%BD%A6%E7%A0%B8%E9%87%91%E8%9B%8B%E5%BE%97%E4%B8%9C%E9%A3%8E%E6%97%A5%E4%BA%A7%E9%99%90%E9%87%8F%E7%89%88%E7%A4%BC%E5%93%81","%E9%99%90%E5%89%8D10%E5%90%8D%E8%AE%A2%E8%BD%A6%E5%AE%A2%E6%88%B7","%E5%87%A1%E8%AE%A2%E8%BD%A6%EF%BC%8C%E5%8D%B3%E5%8F%AF%E5%8F%82%E4%B8%8E5%E8%BD%AE%E5%BE%AA%E7%8E%AF%E6%83%8A%E5%96%9C%E5%A4%A7%E6%8A%BD%E5%A5%96%EF%BC%8C%E5%A4%9A%E9%87%8D%E4%B8%AD%E5%A5%96%E6%9C%BA%E4%BC%9A%EF%BC%9B","%E4%BA%AB%E5%8F%978000%E5%85%83%E8%A1%A5%E8%B4%B4","%E4%BA%AB%E5%8F%974000%E5%85%83%E8%A1%A5%E8%B4%B4","%E4%BA%AB%E5%8F%972000%E5%85%83%E8%A1%A5%E8%B4%B4"]`
		result1.Car_model = `["10060", "10220", "24526", "353861"]`
		result1.Activity_status = "1"
		result1.Createtime = "2015-09-19 11:43:11"
		result1.Updatetime = "2015-09-19 11:43:12"
		result = result1
	} else if id == "4028832b4ec94b0f014ece1595060006" {
		result2 := t_activity{}
		result2.Activity_id = "4028832b4ec94b0f014ece1595060006"
		result2.Activity_type = 1
		result2.Dealer_id = "S2015012829704"
		result2.City_id = 440100
		result2.Title = "2015新楼兰火热预售"
		result2.Title_short = "预订新楼兰 订金赚红利"
		result2.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437977286597.jpg"
		result2.Start_date = "2015-07-29 14:06:22"
		result2.End_date = "2015-08-08 18:00:00"
		result2.Phone = "020-85553008"
		result2.Activity_discount = 0
		result2.Activity_address = "中山大道西895号棠东（骏景花园对面）"
		result2.Detail = `1、预售期间订车，订金享8月8日新车上市当天余额宝的利息红利！
2、订车越早，定金越高，红利越多！
3、所得红利可用于选购新车精品活兑换等值市场礼品！
4、新车上市后，若不满意新车可无条件退订！红利同样可换购礼品！`
		result2.Packages = `["%E6%96%B9%E5%90%91%E7%9B%98%E5%A5%97","%E4%B9%B0%E8%BD%A6%E9%80%81%E8%BD%A6%E8%86%9C"]`
		result2.Car_model = `["353861"]`
		result2.Activity_status = "1"
		result2.Createtime = "2015-09-19 11:43:13"
		result2.Updatetime = "2015-09-19 11:43:14"
		result = result2
	} else if id == "4028832b4ed50893014ed8711f3c0004" {
		result3 := t_activity{}
		result3.Activity_id = "4028832b4ed50893014ed8711f3c0004"
		result3.Activity_type = 1
		result3.Dealer_id = "S2015036970924"
		result3.City_id = 441300
		result3.Title = "奇骏最高优惠1.2万 店内现车充足"
		result3.Title_short = "奇骏直降1.2万"
		result3.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150901/1441089585326.jpg"
		result3.Start_date = "2015-07-29 00:00:00"
		result3.End_date = "2015-08-09 00:00:00"
		result3.Phone = "0752-2285272"
		result3.Activity_discount = 0
		result3.Activity_address = "广东省惠州市三环路与口岸路交接处（实验中学旁）"
		result3.Detail = `本店内奇骏现车销售，颜色可选，目前购车部分加装车型可优惠1.20万元，感兴趣的朋友可以到店咨询购买，咨询电话：0752-2285272`
		result3.Packages = `["%E6%96%B0%E5%A5%87%E9%AA%8F%E7%9B%B4%E9%99%8D1.2%E4%B8%87"]`
		result3.Car_model = `["47710"]`
		result3.Activity_status = "1"
		result3.Createtime = "2015-09-19 11:43:15"
		result3.Updatetime = "2015-09-19 11:43:16"
		result = result3
	} else if id == "4028832b4ed50893014ed87aa5890006" {
		result4 := t_activity{}
		result4.Activity_id = "4028832b4ed50893014ed87aa5890006"
		result4.Activity_type = 1
		result4.Dealer_id = "S2015036970924"
		result4.City_id = 441300
		result4.Title = "天籁现车充足 最高现金优惠1.9万元"
		result4.Title_short = "天籁直降1.9万"
		result4.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437976084376.jpg"
		result4.Start_date = "2015-07-30 00:00:00"
		result4.End_date = "2015-08-30 00:00:00"
		result4.Phone = "0752-2285272"
		result4.Activity_discount = 0
		result4.Activity_address = "广东省惠州市三环路与口岸路交接处（实验中学旁）"
		result4.Detail = `店内天籁现车销售，颜色可选，目前购车部分车型可优惠3.30万元，感兴趣的朋友可以到店咨询购买，热线4008300451`
		result4.Packages = `["%E5%A4%A9%E7%B1%81%E6%9C%80%E9%AB%98%E7%8E%B0%E9%87%91%E4%BC%98%E6%83%A01.9%E4%B8%87%E5%85%83"]`
		result4.Car_model = `["10220", "353861", "47710", "49291", "492911"]`
		result4.Activity_status = "1"
		result4.Createtime = "2015-09-19 11:43:17"
		result4.Updatetime = "2015-09-19 11:43:18"
		result = result4
	} else if id == "4028832b4ef205c5014ef27cd4190000" {
		result5 := t_activity{}
		result5.Activity_id = "4028832b4ef205c5014ef27cd4190000"
		result5.Activity_type = 1
		result5.Dealer_id = "S2015012829704"
		result5.City_id = 440100
		result5.Title = "盛世经典 东风日产“千人限时特卖会”"
		result5.Title_short = "日产“千人限时特卖会”"
		result5.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437976084376.jpg"
		result5.Start_date = "2015-08-09 09:00:00"
		result5.End_date = "2015-08-09 18:00:00"
		result5.Phone = "020-85553008"
		result5.Activity_discount = 0
		result5.Activity_address = "广州市天河区中山大道西895号"
		result5.Detail = `第一重礼：预交诚意金可变3000元!
第二重礼：签到即送价值199大礼包!
第三重礼：新楼兰广州上市，前10名现场订楼兰，享全国优先提车!
第四重礼：现金一把抓，手有多大钱就有多少!
第五重礼：玩转大富翁，大奖任你赢!
第六重礼：精品一元拍，价格者得!
第七重礼：厂家赞助礼—终身延保，排队领取，先到先得，数量有限，送完即止!
第八重礼：总经理特别礼!
第九重礼：前30名订车客户额外获赠价值280元行车记录仪一台!
第十重礼：15周年庆典十轮狂欢礼，十轮大奖疯狂抽取，越早订车中奖越多!`
		result5.Packages = `["%E7%8E%B0%E5%9C%A8%E6%8A%A5%E5%90%8D%E5%8F%82%E5%8A%A0%E6%B4%BB%E5%8A%A8%E6%8A%A2%E5%85%88%E8%B5%A2%E5%8F%961000%E5%85%83%E6%B2%B9%E5%8D%A1%21"]`
		result5.Car_model = `["10060", "10140", "10220", "14660", "24526", "30606", "47710", "49190", "49291", "492911"]`
		result5.Activity_status = "1"
		result5.Createtime = "2015-09-19 11:43:19"
		result5.Updatetime = "2015-09-19 11:43:20"
		result = result5
	} else if id == "4028832b4ef92959014efd1596880004" {
		result6 := t_activity{}
		result6.Activity_id = "4028832b4ef92959014efd1596880004"
		result6.Activity_type = 1
		result6.Dealer_id = "S2015072210924"
		result6.City_id = 370700
		result6.Title = "青山启辰巅峰钜惠 八月送清凉"
		result6.Title_short = "青山启辰巅峰钜惠"
		result6.Picture = "http://www.chebaba.com/userfiles/store/activityModelImg/20150727/1437967984964.jpg"
		result6.Start_date = "2015-08-08 15:00:00"
		result6.End_date = "2015-08-09 23:55:00"
		result6.Phone = "0536-8369002"
		result6.Activity_discount = 0
		result6.Activity_address = "山东省潍坊市潍城区殷大路与北宫西街交叉口300米（潍城区政府以西3公里处）"
		result6.Detail = `青江日产8月8日巅峰购车第四季闪亮登场！东风日产的粉丝们！这是你们享受低价的购车好机会！假如您在我们之前的活动中没考虑好，又或者是在之前的优惠中犹豫了，再或者是一不小心错过了，想买但是又找不到合适的机会，没关系，东风日产潍坊青江专营店8月8日钜惠来袭！狂购盛宴、巅峰钜惠，满足您便捷、低价购车的需求！您无需周旋于各大4S店，无需费尽力气讨价还价，无需提出N种买送要求，8月8日青江日产，您想要的，我们全都给你！`
		result6.Packages = `["%E8%BF%9B%E5%BA%97%E7%AD%BE%E5%88%B0%E5%B0%B1%E5%8F%AF%E5%8F%82%E4%B8%8E%E5%A4%A7%E5%AF%8C%E7%BF%81%E6%B8%B8%E6%88%8F%EF%BC%8C%E7%96%AF%E7%8B%82%E9%80%81%E8%B1%AA%E7%A4%BC%E5%85%88%E5%88%B0%E5%85%88%E5%BE%97%EF%BC%9B","%E7%8E%B0%E5%9C%BA%E8%AE%A2%E8%BD%A6%E5%B0%B1%E9%80%81%E5%8D%83%E5%85%83%E8%BD%A6%E8%A3%85%E5%A4%A7%E7%A4%BC%E5%8C%85%EF%BC%9B","%E7%8E%B0%E5%9C%BA%E8%AE%A2%E8%BD%A6%E5%AE%A2%E6%88%B7%E5%9D%87%E5%8F%AF%E5%8F%82%E4%B8%8E%E7%A0%B8%E9%87%91%E8%9B%8B%E8%B5%A2%E5%A5%BD%E7%A4%BC%EF%BC%8C%E7%B2%BE%E7%BE%8E%E7%A4%BC%E5%93%81%E9%9A%8F%E5%BF%83%E6%8A%BD","%E8%B4%AD%E8%BD%A6%E5%8D%B3%E6%9C%89%E6%9C%BA%E4%BC%9A%E8%8E%B7%E5%BE%971800%E5%85%83%E7%BE%8E%E7%9A%84%E9%AB%98%E7%AB%AF%E5%AE%B6%E7%94%B5%E5%9B%9B%E4%BB%B6%E5%A5%97%EF%BC%9B","%E7%8E%B0%E5%9C%BA%E8%AE%A2%E8%BD%A6%E5%AE%A2%E6%88%B7%E9%99%A4%E4%BB%A5%E4%B8%8A%E7%A4%BC%E5%93%81%E5%A4%96%EF%BC%8C%E8%BF%98%E5%8F%AF%E8%8E%B7%E5%BE%97%E7%BA%AA%E5%BF%B5%E7%A4%BC%E5%93%81%E4%B8%80%E4%BB%BD%EF%BC%9B","%E8%80%81%E5%AE%A2%E6%88%B7%E6%8E%A8%E8%8D%90%E6%96%B0%E5%AE%A2%E6%88%B7%E6%88%90%E5%8A%9F%E8%B4%AD%E8%BD%A6%EF%BC%8C%E5%8D%B3%E9%80%81%E7%9B%B8%E5%BA%94%E7%A7%AF%E5%88%86%EF%BC%8C%E7%A7%AF%E5%88%86%E5%BD%93%E9%92%B1","%E8%BF%9B%E5%BA%97%E8%AF%84%E4%BC%B0%E4%BA%8C%E6%89%8B%E8%BD%A6%E3%80%81%E5%9B%9E%E5%8E%82%E4%BF%9D%E5%85%BB%E7%88%B1%E8%BD%A6%EF%BC%8C%E7%B2%BE%E7%BE%8E%E7%A4%BC%E5%93%81%E9%9A%8F%E5%BF%83%E4%BA%AB%EF%BC%81"]`
		result6.Car_model = `["43013", "430131", "430132", "49299", "49300"]`
		result6.Activity_status = "1"
		result6.Createtime = "2015-09-19 11:43:21"
		result6.Updatetime = "2015-09-19 11:43:22"
		result = result6
	} else {
		c.JSON(404, gin.H{"result": nil})
		return
	}

	c.JSON(200, gin.H{"result": result})
}
