package integration

import (
	"os/exec"
	"strconv"
	"strings"

	"code.cloudfoundry.org/grootfs/groot"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func CreateBundle(grootFSBin, graphPath, imagePath, id string) groot.Bundle {
	cmd := exec.Command(grootFSBin, "--graph", graphPath, "create", "--image", imagePath, id)
	sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(sess).Should(gexec.Exit(0))

	return groot.NewBundle(strings.TrimSpace(string(sess.Out.Contents())))
}

func DeleteBundle(grootFSBin, graphPath, id string) string {
	cmd := exec.Command(grootFSBin, "--graph", graphPath, "delete", id)
	sess, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(sess).Should(gexec.Exit(0))
	return string(sess.Out.Contents())
}

func FindUID(user string) uint32 {
	sess, err := gexec.Start(exec.Command("id", "-u", user), nil, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(sess).Should(gexec.Exit(0))

	i, err := strconv.ParseInt(strings.TrimSpace(string(sess.Out.Contents())), 10, 32)
	Expect(err).NotTo(HaveOccurred())

	return uint32(i)
}

func FindGID(group string) uint32 {
	sess, err := gexec.Start(exec.Command("id", "-g", group), nil, GinkgoWriter)
	Expect(err).NotTo(HaveOccurred())
	Eventually(sess).Should(gexec.Exit(0))

	i, err := strconv.ParseInt(strings.TrimSpace(string(sess.Out.Contents())), 10, 32)
	Expect(err).NotTo(HaveOccurred())

	return uint32(i)
}
