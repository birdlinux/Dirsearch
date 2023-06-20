
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/birdlinux/Subsearch">
    <img src="https://github.com/birdlinux/Subsearch/assets/123122904/403b874b-bad0-4437-907d-06cacc83a8a1" alt="Logo" width="100" height="100">
  </a>

  <h3 align="center">Dirsearch</h3>
 
  <p align="center">
    Enhance subdomain analysis with Dirsearch's effectiveness
    <br />
    <a href="https://github.com/birdlinux/Dirsearch"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/birdlinux/Dirsearch/">View Demo</a>
    ·
    <a href="https://github.com/birdlinux/Dirsearch/issues">Report Bug</a>
    ·
    <a href="https://github.com/birdlinux/Dirsearch/issues">Request Feature</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about">About The Project</a>
      <ul>
        <li><a href="#builtwith">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#installation">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<br />
<center> <h1 align="left" id="about">About</h1> </center>

Dirsearch is an innovative project designed to revolutionize DNS enumeration, providing a powerful tool for effortlessly discovering and mapping subdomains. With its advanced capabilities and streamlined efficiency, Dirsearch simplifies the process of subdomain analysis, offering comprehensive visibility into hidden domains.

![image](https://github.com/birdlinux/Dirsearch/assets/123122904/d02deca4-9827-4760-8ba6-6b5d4a68c93a)



<br />
<center> <h1 align="left" id="builtwith">Built with</h1> </center>

* [go.dev](https://go.dev/)

<br />
<center> <h1 align="left" id="installation">Installation</h1> </center>

_Below is an example of how you can instruct your audience on installing and setting up your binary file. This template doesn't rely on any external services._

1. Clone the repo
   ```sh
   git clone https://github.com/birdlinux/Dirsearch.git
   ```
2. Go into the Directory
   ```sh
   cd Dirsearch
   ```

3. Building the binary
   ```sh
   go build release/dirsearch.go
   ```
   
4. Move to /usr/bin
    ```sh
    sudo chmod +x dirsearch
    ```

    ```sh
    sudo mv dirsearch /usr/bin/
    ```
    
<br />
<center> <h1 align="left" id="usage">Usage</h1> </center>

1. Run Dirsearch
```sh
dirsearch --domain https://example.com --wordlist folders.txt
```

See the [open issues](https://github.com/birdlinux/Dirsearch/issues) for a full list of proposed features (and known issues).

<br />
<center> <h1 align="left" id="contributing">Contributing</h1> </center>

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b new-dev-20220608`)
3. Commit your Changes (`git commit -m 'Added Language Support'`)
4. Push to the Branch (`git push origin new-dev-20220608`)
5. Open a Pull Request


<!-- LICENSE -->
<br />
<center> <h1 align="left" id="license">License</h1> </center>

Distributed under the MIT License. See `LICENSE` for more information.
