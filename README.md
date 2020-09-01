 使用示例
 ```bash
 grafana-manager --url "http://172.31.107.24:38082/" -u admin -p admin import dashboard ./test/test-dashboard
 grafana-manager import datasource ./test/test-datasource --url "http://172.31.107.24:38082/" -u admin -p admin 
```
 
错误:

json: cannot unmarshal string into Go struct field Board.panels of type int

解决方法：

查看对应的json文件min_doc_count这个字段是不是为string，改为int类型

如： "min_doc_count": "1" 改为 "min_doc_count": 1
