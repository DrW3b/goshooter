# GoShooter
![Go Logo](https://camo.githubusercontent.com/b864130864173a91916143250a96a36effd3752914b3d678607842a2ca56def2/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f476f2d3030414444383f7374796c653d666f722d7468652d6261646765266c6f676f3d676f266c6f676f436f6c6f723d7768697465)
![bash logo](https://camo.githubusercontent.com/aca8077e4bfa77bc5469b4691a9f649a1e22ea5a3271f82bb09dbc7cff80bf4c/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f5368656c6c5f5363726970742d3132313031313f7374796c653d666f722d7468652d6261646765266c6f676f3d676e752d62617368266c6f676f436f6c6f723d7768697465)

GoShooter is a CLI tool written in Go for taking screenshots of websites provided in a list.

## Prerequisites

You should have the following installed on your system before proceeding:

- Go (version 1.15 or higher)
- Google Chrome

## Installation

To install GoShooter on your system, follow these steps:

1. Install the required Go package and disable Go modules:

    ```bash
    go env -w GO111MODULE=off
    go get -u github.com/chromedp/chromedp
    ```

2. Clone the GoShooter repository:

    ```bash
    git clone https://github.com/DrW3b/goshooter.git
    ```

3. Navigate to the GoShooter directory:

    ```bash
    cd goshooter
    ```


4. Run the installation script to build and move `goshooter` to `/usr/local/bin`:

    ```bash
    chmod +x install.sh
    ./install.sh
    ```

Once you've completed these steps, you should be able to run the `goshooter` command from any location in your terminal or command prompt and you can remove the folder goshooter.

## Usage

To use GoShooter, follow these steps:

1. Create a file containing the URLs you want to capture screenshots for. The URLs should be domain names, without "http://" or "https://". For example, your file (`domains.txt`) might look like this:

    ```bash
    google.com
    yahoo.com
    bing.com
    ```

2. Run the GoShooter command:

    ```bash
    goshooter -f domains.txt -t 5s -th 10
    ```

This command tells GoShooter to capture screenshots for the domains listed in `domains.txt`, using 10 threads and a 5-second timeout.

## Disclaimer

This software is provided by DrW3B "as is" and without any warranties. DrW3B is not liable for any damages or consequences that may occur due to the use or misuse of this software.


