// Package scan provides types and function to perform TCP port scans on a list of hosts
package scan

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
)

type HostsList struct {
	Hosts []string
}

// Searches for hosts in the list
func (hl *HostsList) search(host string) (bool, int) {
	sort.Slice(hl.Hosts, func(i, j int) bool {
		return hl.Hosts[i] < hl.Hosts[j]
	})

	i := sort.SearchStrings(hl.Hosts, host)
	if i < len(hl.Hosts) && hl.Hosts[i] == host {
		return true, i
	}

	return false, -1
}

// Adds a host to the list
func (hl *HostsList) Add(host string) error {
	if found, _ := hl.search(host); found {
		return fmt.Errorf("Host %s already in the list", host)
	}

	hl.Hosts = append(hl.Hosts, host)
	return nil
}

// Removes a host from the list
func (hl *HostsList) Remove(host string) error {
	if found, i := hl.search(host); found {
		hl.Hosts = append(hl.Hosts[:i], hl.Hosts[i+1:]...)
		return nil
	}

	return fmt.Errorf("Host %s is not in the list", host)
}

// Obtains hosts from a hosts file
func (hl *HostsList) Load(hostsFile string) error {
	f, err := os.Open(hostsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		hl.Hosts = append(hl.Hosts, scanner.Text())
	}

	return nil
}

// Saves hosts to a hosts file
func (hl *HostsList) Save(hostsFile string) error {
	output := ""

	for _, h := range hl.Hosts {
		output += fmt.Sprintln(h)
	}

	return ioutil.WriteFile(hostsFile, []byte(output), 0644)
}
