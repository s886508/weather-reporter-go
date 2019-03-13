
# Introduction
This is a simple crawler built by Go. Input city to lookup weather data from next few days.

# Start
- Run by interpreter
`go run main`
- Build binary
`go install main`

# Example
The tool run until user enter `Q`. Enter`city name` and `days` to print retrieved data.
```
Please enter city name(Q to exit): �O�_��
Please enter days to show(1~7): 5
============ �O�_�� ============

 03/13 �P���T     �դ�:15 ~ 22�XC:�h��
 03/13 �P���T     �ߤW:17 ~ 20�XC:���ɦh���u�Ȱ}�B

 03/14 �P���|     �դ�:17 ~ 21�XC:���u�Ȱ}�B
 03/14 �P���|     �ߤW:16 ~ 19�XC:�h���ɳ��u�Ȱ}�B

 03/15 �P����     �դ�:16 ~ 19�XC:�h���ɳ��u�ȫB
 03/15 �P����     �ߤW:16 ~ 17�XC:�h���u�ȫB

 03/16 �P����     �դ�:16 ~ 19�XC:���u�ȫB
 03/16 �P����     �ߤW:16 ~ 18�XC:���ɦh���u�ȫB

 03/17 �P����     �դ�:16 ~ 19�XC:�h���ɳ��u�Ȱ}�B
 03/17 �P����     �ߤW:15 ~ 17�XC:���u�Ȱ}�B

============ �O�_�� ============
```
