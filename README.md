# documents
유용한 정보 모음
# Installing Git with Default Packages
Debian’s default repositories provide you with a fast method to install Git. Note that the version you install via these repositories may be older than the newest version currently available. If you need the latest release, consider moving to the next section of this tutorial to learn how to install and compile Git from source.

First, use the apt package management tools to update your local package index. With the update complete, you can download and install Git:  
'''
$sudo apt update
$sudo apt install git
'''  
You can confirm that you have installed Git correctly by running the following command:

g$it --version
Output
git version 2.11.0

With Git successfully installed, you can now move on to the Setting Up Git section of this tutorial to complete your setup.

#Installing Git from Source
A more flexible method of installing Git is to compile the software from source. This takes longer and will not be maintained through your package manager, but it will allow you to download the latest release and will give you some control over the options you include if you wish to customize.

Before you begin, you need to install the software that Git depends on. This is all available in the default repositories, so we can update our local package index and then install the packages.

sudo apt update
sudo apt install make libssl-dev libghc-zlib-dev libcurl4-gnutls-dev libexpat1-dev gettext unzip
After you have installed the necessary dependencies, you can go ahead and get the version of Git you want by visiting the Git project’s mirror on GitHub, available via the following URL:

https://github.com/git/git
From here, be sure that you are on the master branch. Click on the Tags link and select your desired Git version. Unless you have a reason for downloading a release candidate (marked as rc) version, try to avoid these as they may be unstable.

git change branch select tags

Next, on the right side of the page, click on the Clone or download button, then right-click on Download ZIP and copy the link address that ends in .zip.

right-click on download zip to copy url

Back on your Debian 9 server, move into the tmp directory to download temporary files.

cd /tmp
From there, you can use the wget command to install the copied zip file link. We’ll specify a new name for the file: git.zip.

wget https://github.com/git/git/archive/v2.18.0.zip -O git.zip
Unzip the file that you downloaded and move into the resulting directory by typing:

unzip git.zip
cd git-*
Now, you can make the package and install it by typing these two commands:

make prefix=/usr/local all
sudo make prefix=/usr/local install
To ensure that the install was successful, you can type git --version and you should receive relevant output that specifies the current installed version of Git.

Now that you have Git installed, if you want to upgrade to a later version, you can clone the repository, and then build and install. To find the URL to use for the clone operation, navigate to the branch or tag that you want on the project’s GitHub page and then copy the clone URL on the right side:

git copy URL

At the time of writing, the relevant URL is:

https://github.com/git/git.git
Change to your home directory, and use git clone on the URL you just copied:

cd ~
git clone https://github.com/git/git.git
This will create a new directory within your current directory where you can rebuild the package and reinstall the newer version, just like you did above. This will overwrite your older version with the new version:

cd git
make prefix=/usr/local all
sudo make prefix=/usr/local install
With this complete, you can be sure that your version of Git is up to date.

