## Setup Ubuntu
- sudo apt purge --auto-remove nodejs npm
- bcoz the apt is very slow for nodejs
- install nodejs using .tar.xz from website
- extract it and then,
- sudo cp -r ./node-v20.10.0-linux-x64/{bin,include,lib,share} /usr/
- then try new terminal and type node --version

## Using Yarn
- yarn add expo
- yarn expo install? or npx expo install
- yarn expo start

## To Build with EAS
- install android studio & dependencies || https://developer.android.com/studio/install#linux
- sudo apt install libc6:i386 libncurses5:i386 libstdc++6:i386 lib32z1 libbz2-1.0:i386
- after extract:
- sudo cp -r ./android-studio /usr/local/
- then, inside usr/local/android-studio/bin
- execute: ./studio.sh
- make sure to setup sdk path at eas.json ==> mypreview.env.ANDROID_SDK_ROOT: "XXX" (open android studio to check the SDK path)
- yarn global add eas-cli XXXX WRONG
- - Not recommended to use yarn to install global
- - instead, use npm install -g eas-cli
- eas whoami
- eas login (if not yet)
- eas build:configure
- eas build -p android --profile mypreview --local
- still failed in the end lol