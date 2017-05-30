<h1>Dragon-go</h1></br>
Application that fights vs the knights of the kingdom! </br>
Api documentation -> http://www.dragonsofmugloar.com/</br>
<h2>How to run</h2>
Go to project directory & build it by using <b>go build</b> command</br>
Then execute application by lounching the binary file -> <b>./dragon-go</b> (in linux case)</br>
<h2>Branches</h2>
<ul>
  <li>master</li>
  <li>logging</li>
</ul>
In logging branch you can find implemented solution with logging possabilities (needs some tweaks regarding log loss)</br>
<h2>Tests</h2>
Testify was used for covering application logic with tests.</br>
To run tests, go to source directory and type command <b>go test -timeout 30s -tags  ./...</b></br></br>
<b>Package coverage (master branch):</b></br>
<ul>
<li>main - 28%</li>
<li>dragon - 100%</li>
<li>handlers - 77%</li>
<li>result - 100%</li>
<li>weather - 100%</li>
</ul>

<h2>Dependencies:</h2>
Please add dependencies by using guides listed below</br>
Logging -> uber zap (https://github.com/uber-go/zap#installation)</br>
Testing -> testify (https://github.com/stretchr/testify#installation)
