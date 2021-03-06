package task

import (
	typ "github.com/Xuanwo/storage/types"

	"github.com/qingstor/noah/utils"
)

func (t *MoveDirTask) new() {}
func (t *MoveDirTask) run() {
	x := NewListDir(t)
	utils.ChooseSourceStorage(x, t)
	x.SetFileFunc(func(o *typ.Object) {
		sf := NewMoveFile(t)
		sf.SetSourcePath(o.Name)
		sf.SetDestinationPath(o.Name)
		t.GetScheduler().Async(sf)
	})
	x.SetDirFunc(func(o *typ.Object) {
		sf := NewMoveDir(t)
		sf.SetSourcePath(o.Name)
		sf.SetDestinationPath(o.Name)
		t.GetScheduler().Sync(sf)
	})
	t.GetScheduler().Sync(x)
}

func (t *MoveFileTask) new() {}
func (t *MoveFileTask) run() {
	ct := NewCopyFile(t)
	ct.SetCheckTasks(nil)
	t.GetScheduler().Sync(ct)

	dt := NewDeleteFile(t)
	utils.ChooseSourceStorage(dt, t)
	t.GetScheduler().Sync(dt)
}
