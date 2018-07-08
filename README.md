# pdf2img
#基于 ImageMagick-7.0.8 + Ghostscript 9.23  
将pdf文件转为图片  

#Install  

##window(64bits)  
Thanks @vprus
1. 安装 [msys2-x86_64](http://www.msys2.org/)  ps:注意32位和64位的区别  
2. 运行 C:/msys64/msys2.exe  
3. 在msys2 命令行依次输入下列命令 (详细的msys命令见说明[MSYS2 installation](https://github.com/msys2/msys2/wiki/MSYS2-installation))  
```
pacman -Syuu  
pacman -S mingw-w64-x86_64-gcc  
pacman -S mingw-w64-x86_64-pkg-config  
pacman -S mingw-w64-x86_64-zlib  
pacman -S mingw-w64-x86_64-imagemagick  
pacman -S mingw-w64-x86_64-ghostscript  #添加对pdf文件的支持  
```
4. 设置环境变量  
4.1 设置临时变量（当前cmd命令窗口有效）  
在cmd窗口里面依次执行下面的命令  
```
PATH=<msys64>\mingw64\bin;%PATH%  
PKG_CONFIG_PATH=<msys64>\mingw64\lib\pkgconfig  
MAGICK_CODER_MODULE_PATH=<msys64>\mingw64\lib\ImageMagick-$VERSION\modules-Q16HDRI\coders  
```
4.2 设置系统变量  <msys64> =msys64安装路径 （eg：C:\msys64）  
我的电脑=>属性=>高级系统设置=>环境变量  
依次添加/编辑下面几个变量  
PATH新增<msys64>\mingw64\bin;  
新增变量PKG_CONFIG_PATH=<msys64>\mingw64\lib\pkgconfig  
新增变量MAGICK_CODER_MODULE_PATH=<msys64>\mingw64\lib\ImageMagick-$VERSION\modules-Q16HDRI\coders  


##centos  
sudo yum install  


##Mac OS X  
MacPorts  
sudo port install ImageMagick  


##Ubuntu / Debian  
sudo apt-get install libmagickwand-dev  


5. 检测是 ImageMagick 是否正确配置  
在cmd里面运行命令  
```
pkg-config --cflags --libs MagickWand  
```
如果返回结果类似下面的结果 说明配置成功  
```
-fopenmp -DMAGICKCORE_HDRI_ENABLE=1 -DMAGICKCORE_QUANTUM_DEPTH=16 -D_DLL -D_MT -IC:/msys64/mingw64/include/ImageMagick-7  -LC:/msys64/mingw64/lib -lMagickWand-7.Q16HDRI -lMagickCore-7.Q16HDRI  
```
6. 下载imagemagick官方推荐库(golang)[imagick](https://github.com/gographics/imagick)  
在cmd里面运行 
```
go get gopkg.in/gographics/imagick.v3/imagick  
```

ps: 库的版本和imagemagick版本的对应关系如下  
| 分支/tag|ImageMagick版本范围|对应gopkg.in|  
|legacy (tag v1.x.x)| 6.7.x   <= ImageMagick <= 6.8.9-10 |gopkg.in/gographics/imagick.v1/imagick|  
|master (tag v2.x.x)| 6.9.0-2 <= ImageMagick <= 6.9.9-35 |gopkg.in/gographics/imagick.v2/imagick|  
|im-7   (tag v3.x.x)| 7.x     <= ImageMagick <= 7.x |gopkg.in/gographics/imagick.v3/imagick|  


#WORK  
在imagick的基础上进行自定义开发  

#PS  
1. imagemagick官方地址[imagemagick](https://www.imagemagick.org/)  
2. ghostscript官方地址[ghostscript](https://ghostscript.com/download/gsdnld.html)  
