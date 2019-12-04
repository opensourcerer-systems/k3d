/*
Copyright © 2019 Thorsten Klein <iwilltry42@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cluster

import (
	"github.com/rancher/k3d/pkg/runtimes"
	k3d "github.com/rancher/k3d/pkg/types"
	log "github.com/sirupsen/logrus"
)

// LoadImagesIntoCluster starts up a k3d tools container for the selected clusters and uses it to export
// images from the runtime to import them into the nodes of the selected cluster
func LoadImagesIntoCluster(runtime runtimes.Runtime, images []string, clusters *k3d.Cluster) {

}

// startToolsContainer will start a new k3d tools container and connect it to the network of the chosen cluster
func startToolsContainer(runtime runtimes.Runtime, cluster *k3d.Cluster) (string, error) {
	log.Debugln("starttoolscontainer")

	return "", nil
}
