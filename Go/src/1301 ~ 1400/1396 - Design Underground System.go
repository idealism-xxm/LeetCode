// 链接：https://leetcode.com/problems/design-underground-system/
// 题意：实现一个能计算地铁平均耗时的数据结构，支持查询不同出入站的用户平均耗时。
//
//      该数据结构需要支持以下操作：
//          1. void checkIn(int id, string stationName, int t): 
//              用户 id 在时间 t 从 stationName 进入地铁
//          2. void checkOut(int id, string stationName, int t): 
//              用户 id 在时间 t 从 stationName 离开地铁
//          3. double getAverageTime(string startStation, string endStation):
//              返回从 startStation 到 endStation 的所有用户的平均耗时
//
//      一个用户的的入站和出栈事件必定是合法的，且入站时间小于出站时间，
//      并且所有的事件按照时间升序发生。


// 数据限制：
//  1 <= id, t <= 10 ^ 6
//  1 <= stationName.length, startStation.length, endStation.length <= 10
//  所有的字符串均由英文大小写字母和数字组成
//  最多会总共调用 2 * 10 ^ 4 次 checkIn, checkOut, getAverageTime 方法
//  答案需要精确到小数点后 5 位


// 输入： ["UndergroundSystem","checkIn","checkIn","checkIn","checkOut","checkOut","checkOut","getAverageTime","getAverageTime","checkIn","getAverageTime","checkOut","getAverageTime"]
//       [[],[45,"Leyton",3],[32,"Paradise",8],[27,"Leyton",10],[45,"Waterloo",15],[27,"Waterloo",20],[32,"Cambridge",22],["Paradise","Cambridge"],["Leyton","Waterloo"],[10,"Leyton",24],["Leyton","Waterloo"],[10,"Waterloo",38],["Leyton","Waterloo"]]
// 输出： [null,null,null,null,null,null,null,14.00000,11.00000,null,11.00000,null,12.00000]
// 解释： UndergroundSystem undergroundSystem = new UndergroundSystem();
//       undergroundSystem.checkIn(45, "Leyton", 3);
//       undergroundSystem.checkIn(32, "Paradise", 8);
//       undergroundSystem.checkIn(27, "Leyton", 10);
//       undergroundSystem.checkOut(45, "Waterloo", 15);  // 用户 45 出入站为 "Leyton" -> "Waterloo" ，耗时为 15-3 = 12
//       undergroundSystem.checkOut(27, "Waterloo", 20);  // 用户 27 出入站为 "Leyton" -> "Waterloo" ，耗时为 20-10 = 10
//       undergroundSystem.checkOut(32, "Cambridge", 22); // 用户 32 出入站为 "Paradise" -> "Cambridge" ，耗时为 22-8 = 14
//       undergroundSystem.getAverageTime("Paradise", "Cambridge"); // 返回 14.00000 。 有 1 个用户出入站为 "Paradise" -> "Cambridge" ， (14) / 1 = 14
//       undergroundSystem.getAverageTime("Leyton", "Waterloo");    // 返回 11.00000 。 有 2 个用户出入站为 "Leyton" -> "Waterloo" ， (10 + 12) / 2 = 11
//       undergroundSystem.checkIn(10, "Leyton", 24);
//       undergroundSystem.getAverageTime("Leyton", "Waterloo");    // 返回 11.00000
//       undergroundSystem.checkOut(10, "Waterloo", 38);  // 用户 10 出入站为 "Leyton" -> "Waterloo" ，耗时为 38-24 = 14
//       undergroundSystem.getAverageTime("Leyton", "Waterloo");    // 返回 12.00000. 有 3 个用户出入站为 "Leyton" -> "Waterloo", (10 + 12 + 14) / 3 = 12

// 输入： ["UndergroundSystem","checkIn","checkOut","getAverageTime","checkIn","checkOut","getAverageTime","checkIn","checkOut","getAverageTime"]
//       [[],[10,"Leyton",3],[10,"Paradise",8],["Leyton","Paradise"],[5,"Leyton",10],[5,"Paradise",16],["Leyton","Paradise"],[2,"Leyton",21],[2,"Paradise",30],["Leyton","Paradise"]]
// 输出： [null,null,null,5.00000,null,null,5.50000,null,null,6.66667]
// 解释： UndergroundSystem undergroundSystem = new UndergroundSystem();
//       undergroundSystem.checkIn(10, "Leyton", 3);
//       undergroundSystem.checkOut(10, "Paradise", 8); // 用户 10 出入站为 "Leyton" -> "Paradise" ，耗时为 8-3 = 5
//       undergroundSystem.getAverageTime("Leyton", "Paradise"); // 返回 5.00000 ， (5) / 1 = 5
//       undergroundSystem.checkIn(5, "Leyton", 10);
//       undergroundSystem.checkOut(5, "Paradise", 16); // 用户 5 出入站为 "Leyton" -> "Paradise"  16-10 = 6
//       undergroundSystem.getAverageTime("Leyton", "Paradise"); // 返回 5.50000 ， (5 + 6) / 2 = 5.5
//       undergroundSystem.checkIn(2, "Leyton", 21);
//       undergroundSystem.checkOut(2, "Paradise", 30); // 用户 2 出入站为 "Leyton" -> "Paradise" ，耗时为 30-21 = 9
//       undergroundSystem.getAverageTime("Leyton", "Paradise"); // 返回 6.66667 ， (5 + 6 + 9) / 3 = 6.66667

// 思路： Map
//
//      可以使用两个 map 来维护所需信息：
//          1. id_to_start_time: 维护仅入站用户 id 的入站名和入站时间
//          2. id_to_end_time:
//              维护出入站为 (start_station, end_station) 的所有用户的个数与总耗时
//
//
//      然后针对不同的操作进行处理即可：
//          1. checkIn: 将其 id 和入站名和入站时间存入 id_to_start_time 。
//          2. checkOut: 从 id_to_start_time 获取其入站名和入站时间，
//              然后计算该用户出入站耗时，
//              再更新 travel_to_info 出入站的用户数和总耗时即可
//          3. getAverageTime: 从 travel_to_info 获取出入站的用户数和总耗时，
//              计算平均耗时返回即可
//
//      时间复杂度：O(1)
//          1. 三个方法都仅对 map 进行了常数次查询和修改，
//              所以时间复杂度为 O(1)
//      空间复杂度：O(n)
//          1. 需要存储全部 O(n) 个入站信息
//          2. 每次统计​出入站记录信息前，都会删除入站的信息，所以不会增加额外空间消耗


type StartInfo struct {
    // 入站名
    station string
    // 入站时间
    time int
}


type StationInfo struct {
    // 出入站的用户数
    cnt int
    // 出入站的总耗时
    totalTime int
}


type UndergroundSystem struct {
    // 维护仅入站用户 id 的入站名和入站时间
    idToStartTime map[int]StartInfo
    // 维护出入站为 (start_station, end_station) 的所有用户的个数与总耗时
    travelToInfo map[string]StationInfo
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{
        idToStartTime: make(map[int]StartInfo),
        travelToInfo: make(map[string]StationInfo),
    }
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    // 将用户 id 的入站信息放入 id_to_start_time 中
    this.idToStartTime[id] = StartInfo {
        station: stationName,
        time: t,
    }
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    // 获取用户 id 的入站信息
    startInfo := this.idToStartTime[id]
    // 从 id_to_start_time 中删除 id 的入站信息
    delete(this.idToStartTime, id)
    // 获取出入站为 (start_station, end_station) 的所有用户的个数与总耗时信息
    key := fmt.Sprintf("%s-%s", startInfo.station, stationName)
    stationInfo := this.travelToInfo[key]
    // 加上用户 id 的出入站次数
    stationInfo.cnt ++
    // 加上用户 id 的出入站耗时
    stationInfo.totalTime += t - startInfo.time
    // 放回 travelToInfo 中
    this.travelToInfo[key] = stationInfo

}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    // 获取出入站为 (start_station, end_station) 的所有用户的个数与总耗时信息
    key := fmt.Sprintf("%s-%s", startStation, endStation)
    stationInfo := this.travelToInfo[key]
    // 计算平均耗时
    return float64(stationInfo.totalTime) / float64(stationInfo.cnt)
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
