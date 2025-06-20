>   **TSMaster 功能简介**

[**TSMaster 功能简介**	1](#_Toc64824821)

[TSMaster功能特点	3](#_Toc64824822)

>   [1.1 应用管理器	3](#_Toc64824823)

>   [1.2 虚拟通道	3](#_Toc64824824)

>   [1.3 TSMaster小程序	3](#_Toc64824825)

[1.3.1 TSMaster小程序特点	3](#_Toc64824826)

[1.3.2 TSMaster小程序的使用	4](#_Toc64824827)

[1.3.3 小程序的制作	4](#_Toc64824828)

[1.3.4 小程序的IP保护	4](#_Toc64824829)

>   [1.4 TSMaster测试系统	5](#_Toc64824830)

>   [1.5 TSMaster Excel测试模块	5](#_Toc64824831)

>   [1.6 TSMaster自动化服务器	6](#_Toc64824832)

[1.6.1 TSMaster进程内自动化服务器	6](#_Toc64824833)

[1.6.2 TSMaster进程外自动化服务器	6](#_Toc64824834)

>   [1.7 系统变量模块	6](#_Toc64824835)

>   [1.8 符号映射模块	6](#_Toc64824836)

>   [1.9 剩余总线仿真	7](#_Toc64824837)

>   [1.10 仿真面板	7](#_Toc64824838)

>   [1.11 TSMaster图形编程语言	7](#_Toc64824839)

>   [1.12 诊断模块	8](#_Toc64824840)

>   [1.13 标定模块	8](#_Toc64824841)

>   [1.14 激励模块	8](#_Toc64824842)

>   [1.15 参数曲线模块	8](#_Toc64824843)

>   [1.16 数据库功能	8](#_Toc64824844)

>   [1.17 测量功能	9](#_Toc64824845)

[1.17.1 总线Trace功能	9](#_Toc64824846)

[1.17.2 报文发送功能	9](#_Toc64824847)

[1.17.3 图形显示功能	9](#_Toc64824848)

[1.17.4 仪表显示功能	9](#_Toc64824849)

[1.17.5 总线统计功能	9](#_Toc64824850)

>   [1.18 记录功能	9](#_Toc64824851)

>   [1.19 数据回放功能	9](#_Toc64824852)

[1.19.1 在线数据回放	10](#_Toc64824853)

[1.19.2 视频回放	10](#_Toc64824854)

[1.19.3 离线数据回放	10](#_Toc64824855)

>   [1.20 记录文件转换模块	10](#_Toc64824856)

>   [1.21 通道映射模块	10](#_Toc64824857)

>   [1.22 Turbo极速模式	10](#_Toc64824858)

>   [1.23 硬件配置模块	10](#_Toc64824859)

>   [1.24 窗口管理器	11](#_Toc64824860)

[1.24.1 多窗口引擎	11](#_Toc64824861)

[1.24.2 MDI窗口模式	11](#_Toc64824862)

[1.24.3 Normal窗口模式	11](#_Toc64824863)

[1.24.4 多页模式	11](#_Toc64824864)

[1.24.5 磁吸模式	11](#_Toc64824865)

[1.24.6 工程导入和导出	11](#_Toc64824866)

[1.24.7 窗口style切换	12](#_Toc64824867)

>   [1.25 应用程序宿主	12](#_Toc64824868)

>   [1.26 汽车网络文件转换器	12](#_Toc64824869)

>   [1.27 文档功能	13](#_Toc64824870)

>   [1.28 多语言引擎	13](#_Toc64824871)

>   [1.29 SDK软件二次开发包	13](#_Toc64824872)

>   [1.30 J1939模块	13](#_Toc64824873)

>   TSMaster功能特点

1.  应用管理器

TSMaster
自带的应用管理器（AppManager）提出了虚拟应用程序的概念，在用户需求和底层实现之间形成一个路由。用户只需要操作虚拟应用程序实现业务逻辑，虚拟应用程序负责处理用户需求与实际工具的交互。应用管理器将应用程序所依赖的不同类型的功能（CAN、LIN、I/O、AD、DA等等）抽象为应用程序所拥有的逻辑通道，从而使得任何部分支持这些功能的工具（TSCANLIN、TSCANMini、TSLINMini、Vector系列产品、PEAK系列产品、Kvaser系列产品等等），都可以作为TSMaster的底层，为应用程序提供服务。这也使得TSMaster可以处在一个更高的层级，统一管理众多不同类型的工具，而这些工具的物理通道，将在TSMaster软件中互通。

1.  虚拟通道

在一些没有工具硬件的场合下，用户仍希望能回放数据、执行脚本、运行仿真以评估网络性能，或检查问题等，TSMaster提供了虚拟硬件，其包含多种类型功能对应的虚拟通道，使得软件在不依赖硬件的情况下依然可以执行绝大多数功能。

1.  TSMaster小程序

TSMaster小程序是基于TSMaster软件独立执行的功能扩展平台，任何一个小程序都可以使用TSMaster提供的完整API框架以实现各类扩展功能。TSMaster通过小程序实现软件平台的开放性和拓展性，用户可以直接使用其他用户编写的小程序以完成特定功能，也可以自行编写小程序分享至TSMaster小程序平台，供其他用户下载和评估。

1.  TSMaster小程序特点

TSMaster小程序最大的特点是无需安装无需卸载，可在TSMaster软件中启动并直接实现客户的需求。小程序是轻量化和碎片化的，其模式简单，切入点直接，其设计本身的目的就是为了解决一个特定的问题，实现一个特定的功能，让用户在使用的时候完全没有负担，也无需学习过程。

TSMaster的编程语言就是纯粹的C语言，相对于其他的编程语言来说非常的基础，没有额外的学习成本，对于没有接触过C语言的工程师来说，也只要一次性的学习投入。基于TSMaster的C语言，用户可以做任何可能的事情，例如将现有的ECU代码放入TSMaster小程序中执行，做SIL测试，等等。

1.  TSMaster小程序的使用

TSMaster小程序与传统的软件使用流程（下载安装包——安装——启动——关闭——不用时卸载）不同，小程序省去了安装和卸载的步骤，用户通过TSMaster软件平台提供的搜索引擎，随时可搜索并使用适合自己的小程序，触手可及，用完即走。例如用户想通过读取excel内容，并从中提取报文ID列表以通过控制TSMaster将其自动发送至总线，只需要在TSMaster软件平台搜索“excel”关键字即可找到TSMaster的excel小程序，通过拷贝并修改小程序自带的示例，很快就可以满足需求；再比如用户想基于TSMaster的测试系统执行ECU级测试，希望控制现有的示波器测量总线位时间，只需要在TSMaster软件平台搜索“bit
time”关键字即可查找到适合的小程序，配置简单的波特率等参数后，直接启动就可以执行测试。

1.  小程序的制作

TSMaster小程序的生成方式及其简单，利用TSMaster的C Code
Editor，或是Windows平台的任何编程语言，只要导出三个函数，都可以轻松制作出能被TSMaster加载并执行的小程序。

其中绝大多数的小程序可直接通过TSMaster的C Code Editor生成，用户在C Code
Editor中写下的任何脚本，即使是测试脚本，都可被TSMaster软件平台识别并实现软件代码重用。

1.  小程序的IP保护

小程序（mp文件）是不带源码的二进制文件格式，可以有效地保护小程序发布者的知识产权，小程序的发布者可在文件中留下联系方式，并向其使用者提供技术服务，从而形成基于TSMaster软件平台的商业模式。

1.  TSMaster测试系统

TSMaster自带完整的测试系统框架，它包含测试系统登录模块、用户管理模块、测试系统配置模块、DUT配置模块、测试参数配置模块、测试用例配置模块、测试报告配置模块、测试引擎、测试报告自动生成器、测试记录模块等等。用户只需要基于TSMaster即可制作出专业的测试系统，并在任何可能的场景下实现测试功能，其场景包含但不限于如下内容：

-   芯片级一致性测试（物理层、数据链路层）

-   ECU级一致性测试

-   网络测试

-   集成测试

-   整车测试

-   功能测试

-   软件测试

测试系统配置可单独保存成配置文件，并分享给其他用户使用。由于测试系统拥有三级权限，用户需要登录才可基于被赋予的权限执行测试，且脚本均为小程序（带源码或是不带源码），测试系统的知识产权也有了相应的保障。

1.  TSMaster Excel测试模块

TSMaster可读写Excel文件，并提供Excel读写API，用户可在Excel文件中编写测试脚本，只要编写的测试脚本符合预定义的规则，即可被TSMaster自动执行，并可生成Excel版本的测试报告。

此模块实现了不用编写一行代码，即可编写出完整的测试系统。只要基于现有的Excel测试规范稍加改动，即可将其变为可执行的测试脚本，省去了编写测试脚本的时间，提高测试效率。

1.  TSMaster自动化服务器

TSMaster自动化服务器基于Windows的COM自动化技术，向Windows用户公开了各类API功能，用户仅需要通过简单的脚本（excel、VBA、python、MATLAB、C++、C\#、LabVIEW等等）即可控制TSMaster实现任何用户希望的功能。TSMaster自动化服务器分为进程内和进程外服务器两种，这两类服务器的区别在于，进程内服务器不带界面，是纯dll调用，高性能；而进程外服务器是exe调用，在调用过程中，用户可看到TSMaster界面的各种反馈。两套自动化服务器最大的特点是，它们的核心功能API完全一样，用户只需要修改类字符串即可实现两者的无缝切换，这在同类的软件中是绝无仅有的。

1.  TSMaster进程内自动化服务器

TSMaster进程内自动化服务器不显示TSMaster的界面，在启动后存在于宿主的进程空间。

1.  TSMaster进程外自动化服务器

TSMaster进程外自动化服务器包含在TSMaster软件中，在启动后，TSMaster主界面将显示，相对于TSMaster进程内自动化服务器，其提供了窗口管理器、工程导入导出等等更多功能。

1.  系统变量模块

TSMaster系统变量模块包含了系统自定义，以及用户定义的两大变量表，用户可通过各类途径（graphics、面板，小程序等等）观察和修改系统变量的值，绘制变量曲线等，甚至可以通过小程序，将ECU内部的临时变量放到系统变量中，在graphics窗口显示。使用系统变量极大提高了基于TSMaster的开发效率。

1.  符号映射模块

TSMaster支持将总线数据库中的信号映射到系统变量中，当总线上的信号发生变化时，对应的系统变量将同步变化，此模块的优势是涉及到工程文件发生变更（例如替换dbc）后，与映射符号关联的各类小程序，曲线绘制窗口等均不需要进行更新，只需要更新符号映射部分即可。

1.  剩余总线仿真

TSMaster可根据每个通道载入的数据库中的节点、报文和信号自动配置剩余总线仿真引擎，并独立启动剩余总线仿真功能，其在基于V模型的汽车ECU开发过程中，可支持全网所有节点仿真、半实物仿真、纯监听功能。用户可任意修改仿真引擎中的网络、节点、报文和信号的激活情况。

1.  仿真面板

用户可在TSMaster界面中使用仿真面板，自行设计所需要的界面，并可在界面上显示、修改总线信号的值，从而实现与仿真环境的全面交互，满足用户自身的业务逻辑需求。

TSMaster面板的一大特色是面板内集成编辑器，可随时编辑调整面板，同时面板基于Direct
X技术，支持三种显示模式，可随着窗体尺寸变化任意缩放面板内的各类控件元素。

TSMaster仿真面板路线如下：

第一阶段：实现面板的基本功能，信号的显示，修改，系统变量的关联，各类用户逻辑的触发等等

第二阶段：按钮等控件关联c回调函数执行任意逻辑，可替代通用IDE

第三阶段：面板界面可不依赖TSMaster界面，成为独立可执行程序，双击面板文件即可打开并执行

1.  TSMaster图形编程语言

TSMaster图形编程语言是TSMaster软件平台的重要功能拓展，用户可在其中，不需要编写一行代码，只需要简单地拖拽、连线，即可制作任意复杂的逻辑框图，框图可直接运行，支持单步调试，并且所编辑的框图即是TSMaster小程序，拥有TSMaster小程序的一切特性。使用TSMaster图形编程语言，用户不需要具备高深的编程基础即可开发出具备一定工业强度的程序，满足用户自身的业务需求。

1.  诊断模块

TSMaster自带的诊断模块支持odx标准，可作为诊断仪，对ECU进行诊断，也可基于TSMaster的自动化测试系统，对ECU的诊断进行自动化测试。

1.  标定模块

TSMaster自带的标定模块支持CCP、XCP协议，支持A2L文件格式，支持MDF文件格式，可对ECU进行在线、离线标定，标定过程稳定可靠，能最大程度占用总线负载，并对标定的数据进行实时存储。

1.  激励模块

此模块可关联系统变量，通过各种可能的方式激励系统变量，控制系统变量产生期望的变化曲线。若系统变量关联了标定参数，则此标定参数就可手动控制修改其值，或是由预先设定的曲线进行控制。

1.  参数曲线模块

此模块是标定模块的附属模块，用于显示标定变量中Curve、Map等类型的变量曲线或曲面，用户可任意更改曲线或曲面上的点。

1.  数据库功能

TSMaster支持CAN、LIN总线的主流DBC、LDF文件格式，可载入多个数据库文件，实时解析数据库中的信号，支持信号的物理值、十六进制值、十进制值、字符串值的任意读写。

1.  测量功能
    1.  总线Trace功能

可显示来自CAN、LIN总线的报文列表，支持绝对时间、相对时间显示；信号物理值显示，Value
Table显示、帧率显示，帧类型和数据显示、错误帧显示；支持各类过滤功能。

1.  报文发送功能

支持CAN、LIN报文的实时发送，循环发送的报文中，每条报文都可以设置发送周期，以及信号变化曲线，循环发送的报文数量没有上限。

1.  图形显示功能

支持CAN、LIN信号和系统变量的实时曲线绘制，曲线极速缩放，以及曲线对应的信号属性值的实时修改。

1.  仪表显示功能

支持CAN、LIN信号和系统变量的实时物理量在仪表中的显示，仪表类型丰富，包括指针、数码管、温度计等等。

1.  总线统计功能

支持总线负载率、帧率、帧类型、错误帧的统计，同时支持CAN控制器的状态显示和发送接收错误计数的显示。

1.  记录功能

支持总线报文的压缩存储。

1.  数据回放功能

    1.  在线数据回放

支持多个记录文件同时在线回放，回放过程可触发、暂停，回放进度可实时显示。

1.  视频回放

可加载主流视频，同时关联在线数据回放引擎，从而实现视频与blf记录文件同步回放。在回放过程中可随时暂停，拖动回放进度，并在指定的进度上进行继续回放，或是单帧视频回放等。在视频单帧步进过程中，当前帧到下一帧的时间内所有的报文将被发出。

1.  离线数据回放

支持记录文件在离线环境中回放，以分析记录文件中的报文内容。

1.  记录文件转换模块

支持主流的汽车记录文件格式互转。

1.  通道映射模块

支持修改各个应用程序的逻辑通道类型，通道数量，和每个逻辑通道与硬件设备的绑定。

1.  Turbo极速模式

在对报文时间戳精度要求苛刻的情况下，可开启极速模式，TSMaster将会使用全部的CPU资源，提升应用程序报文收发的时间精度。

1.  硬件配置模块

可对连接应用程序的各个硬件的通道进行配置，包含波特率、采样点、工作模式、只听模式和终端电阻接入等。

1.  窗口管理器

    1.  多窗口引擎

TSMaster支持同时开启多个Trace、发送、图形等窗口，每个窗口可配置各自的过滤器，信号显示，报文发送功能。窗口的数量没有上限。

1.  MDI窗口模式

支持子窗体内嵌于TSMaster的主窗口中显示。任何子窗体都可以切换成普通窗口，浮动于TSMaster窗口之外显示。

1.  Normal窗口模式

支持子窗体浮动于TSMaster主窗口之外显示。任何子窗体都可以切换成MDI窗口，内嵌于TSMaster主窗口中显示。

1.  多页模式

TSMaster支持所有主窗体内部的子窗口排列成多页进行显示，而外部的子窗口保持不变。

1.  磁吸模式

开启磁吸模式后，子窗口与子窗口，子窗口与主窗口，以及任何窗口和屏幕边缘等靠近一定距离后会自动吸附，从而方便用户排列窗口。

1.  工程导入和导出

TSMaster中任何窗口的配置，包括测试系统的配置和小程序配置，都可以保存成工程，导出到硬盘上，并分享给其他用户，用户导入后，原TSMaster中的所有配置，包括数据库的配置、窗口布局等都会得到恢复。

1.  窗口style切换

用户可切换成各种类型风格的用户界面，有深色，亮色，基于贴图的界面，也有基于矢量图的界面。

1.  应用程序宿主

TSMaster的Application
Host功能可抓取任意可执行程序的主窗口，并安排其在自身的子窗口中显示，通过此方法可以使用TSMaster灵活的窗口管理功能，分布式管理各类工程软件。

1.  汽车网络文件转换器

汽车网络文件转换器支持两大类功能：

-   各种类型文件互转

-   从dbc生成嵌入式C代码

    其支持的输入和输出文件格式列表如下：

| 输出\\输入 | DBC | ARXML | Xlsx | Xls | DBF | YAML | SYM |
|------------|-----|-------|------|-----|-----|------|-----|
| DBC        |     | ●     | ●    | ●   | ●   | ●    | ●   |
| ARXML      | ●   |       | ●    | ●   | ●   | ●    | ●   |
| Xlsx       | ●   | ●     |      | ●   | ●   | ●    | ●   |
| Xls        | ●   | ●     | ●    |     | ●   | ●    | ●   |
| CSV        | ●   | ●     | ●    | ●   | ●   | ●    | ●   |
| Json       | ●   | ●     | ●    | ●   | ●   | ●    | ●   |
| DBF        | ●   | ●     | ●    | ●   |     | ●    | ●   |
| YAML       | ●   | ●     | ●    | ●   | ●   |      | ●   |
| SYM        | ●   | ●     | ●    | ●   | ●   | ●    |     |
| FIBEX      | ●   | ●     | ●    | ●   | ●   | ●    | ●   |
| C Source   | ●   |       |      |     |     |      |     |

1.  文档功能

用户可在TSMaster工程的文档窗口编写任意个针对此工程的文档说明，文档支持rtf格式，可贴图，可编辑字体和段落格式，从而帮助使用此工程的其他工程人员迅速理解当前工程的相关信息。

1.  多语言引擎

用户可在主界面中自由选择TSMaster的语言，并实时完成语言的切换。TSMaster的每个语言保存在不同的语言文件中，这使得TSMaster可以轻易支持任何类型的语言。

1.  SDK软件二次开发包

用户可基于TSMaster自带的C语言、C\#、python、MATLAB等示例，以及软件API说明文档，自行开发出TSMaster小程序、或是基于TSMaster
DLL、或是基于TSMaster进程内自动化服务器，或是基于TSMaster进程外自动化服务器的软件，以实现任何可能的需求。

1.  J1939模块

数据库支持J1939的dbc载入，若被载入的数据库检测到J1939的属性，所关联的通道将自动支持J1939协议，总线分析界面的各个窗口支持J1939报文和信号的显示和调节。
