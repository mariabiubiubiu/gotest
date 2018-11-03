namespace go  my.test.demo
namespace py  my.test.demo

struct HourseSell{
 1: i32 sid,
 2: string sname,
 3: bool ssell=0,
 4: i64 sprice,
}

const map<string,string> MAPCONSTANT = {'hello':'world', 'goodnight':'moon'}

service ClassMember {
    list<HourseSell> List(1:i64 callTime),
    string Get(1:string name),
    void Set(1: HourseSell s),
    bool IsNameExist(1:string name),
}
