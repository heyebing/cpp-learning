#include <string>
#include <unordered_map>
#include <iostream>
class KVstore{
public:
    bool set(const std::string& key,const std::string&value);
    std::string get(const std::string& key);
    bool del(const std::string& key);
    private:
    std::unordered_map<std::string, std::string> store_;
};

bool KVstore::set(const std::string& key,const std::string&value){
    store_[key] = value;
    return true;
}
std::string KVstore::get(const std::string& key){
    auto it = store_.find(key);
    if(it == store_.end()){
        return "";
    }
    return it->second;
}
bool KVstore::del(const std::string& key){
    auto it = store_.find(key);
    if(it == store_.end()){
        return false;
    }
    store_.erase(it);
    return true;
}
int main() {
    KVstore kvstore;
    kvstore.set("name", "kujo jotaro");
    kvstore.set("age", "18");
    std::cout << kvstore.get("name") << std::endl;
    std::cout << kvstore.get("age") << std::endl;
    kvstore.del("age");
    std::cout << kvstore.get("name") << std::endl;
    std::cout << kvstore.get("age") << std::endl;
    return 0;
}
