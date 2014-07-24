// Test Swift filesystem interface
package swift_test

import (
	"testing"

	"github.com/ncw/rclone/fs"
	"github.com/ncw/rclone/fstest/fstests"
	"github.com/ncw/rclone/swift"
)

func init() {
	fstests.NilObject = fs.Object((*swift.FsObjectSwift)(nil))
	fstests.RemoteName = "TestSwift:"
}

// Generic tests for the Fs
func TestInit(t *testing.T)                  { fstests.TestInit(t) }
func TestFsString(t *testing.T)              { fstests.TestFsString(t) }
func TestFsRmdirEmpty(t *testing.T)          { fstests.TestFsRmdirEmpty(t) }
func TestFsRmdirNotFound(t *testing.T)       { fstests.TestFsRmdirNotFound(t) }
func TestFsMkdir(t *testing.T)               { fstests.TestFsMkdir(t) }
func TestFsListEmpty(t *testing.T)           { fstests.TestFsListEmpty(t) }
func TestFsListDirEmpty(t *testing.T)        { fstests.TestFsListDirEmpty(t) }
func TestFsNewFsObjectNotFound(t *testing.T) { fstests.TestFsNewFsObjectNotFound(t) }
func TestFsPutFile1(t *testing.T)            { fstests.TestFsPutFile1(t) }
func TestFsPutFile2(t *testing.T)            { fstests.TestFsPutFile2(t) }
func TestFsListDirFile2(t *testing.T)        { fstests.TestFsListDirFile2(t) }
func TestFsListFile1(t *testing.T)           { fstests.TestFsListFile1(t) }
func TestFsNewFsObject(t *testing.T)         { fstests.TestFsNewFsObject(t) }
func TestFsListFile1and2(t *testing.T)       { fstests.TestFsListFile1and2(t) }
func TestFsRmdirFull(t *testing.T)           { fstests.TestFsRmdirFull(t) }
func TestFsPrecision(t *testing.T)           { fstests.TestFsPrecision(t) }
func TestObjectString(t *testing.T)          { fstests.TestObjectString(t) }
func TestObjectFs(t *testing.T)              { fstests.TestObjectFs(t) }
func TestObjectRemote(t *testing.T)          { fstests.TestObjectRemote(t) }
func TestObjectMd5sum(t *testing.T)          { fstests.TestObjectMd5sum(t) }
func TestObjectModTime(t *testing.T)         { fstests.TestObjectModTime(t) }
func TestObjectSetModTime(t *testing.T)      { fstests.TestObjectSetModTime(t) }
func TestObjectSize(t *testing.T)            { fstests.TestObjectSize(t) }
func TestObjectOpen(t *testing.T)            { fstests.TestObjectOpen(t) }
func TestObjectUpdate(t *testing.T)          { fstests.TestObjectUpdate(t) }
func TestObjectStorable(t *testing.T)        { fstests.TestObjectStorable(t) }
func TestObjectRemove(t *testing.T)          { fstests.TestObjectRemove(t) }
func TestObjectPurge(t *testing.T)           { fstests.TestObjectPurge(t) }
func TestFinalise(t *testing.T)              { fstests.TestFinalise(t) }
