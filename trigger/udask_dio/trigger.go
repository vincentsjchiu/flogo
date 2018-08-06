package udask_dio
//#cgo LDFLAGS:-L/home/adlink/Desktop/usb-dask_101_x86_64_0522/lib -lusb_dask64
//#cgo CFLAGS:-I/home/adlink/Desktop/usb-dask_101_x86_64_0522/include
//#include "udask.h"
import "C"
import"unsafe"
import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
)

// MyTriggerFactory My Trigger factory
type MyTriggerFactory struct{
	metadata *trigger.Metadata
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &MyTriggerFactory{metadata:md}
}

//New Creates a new trigger instance for a given id
func (t *MyTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	return &MyTrigger{metadata: t.metadata, config:config}
}

// MyTrigger is a stub for your Trigger implementation
type MyTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
}

// Init implements trigger.Trigger.Init
func (t *MyTrigger) Init(runner action.Runner) {
	t.runner = runner
}

// Metadata implements trigger.Trigger.Metadata
func (t *MyTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

// Start implements trigger.Trigger.Start
func (t *MyTrigger) Start() error {
	// start the trigger
        var cardnum,ChanConfig C.U16;
        var card,err C.I16;
        var Value C.U32;
        var Voltage C.F64;
        var Channel,AdRange C.U16;
        cardnum=0;
        Channel=0;
        AdRange = C.AD_B_10_V;
        card=C.UD_Register_Card(C.USB_2405,cardnum)
        fmt.Printf("%d\n",card)
        ChanConfig = ( C.P2405_AI_EnableIEPE | C.P2405_AI_Coupling_AC | C.P2405_AI_Differential);
        fmt.Printf("%d\n",ChanConfig)
        err = C.UD_AI_2405_Chan_Config( C.U16(card), ChanConfig, ChanConfig, ChanConfig, ChanConfig );
        fmt.Printf("%d\n",err)
        i:=0
        for i<500{
        err = C.UD_AI_ReadChannel(C.U16(card), Channel, AdRange,(*C.U16)(unsafe.Pointer(&Value)) );
        err=C.UD_AI_VoltScale32(C.U16(card), AdRange, 0, Value, &Voltage);
        fmt.Printf("%8.6f\n",Voltage)
        i+=1;

        }
        err =C.UD_Release_Card(C.U16(card));
        fmt.Printf("Release Card%d\n",err)

	return nil
}

// Stop implements trigger.Trigger.Start
func (t *MyTrigger) Stop() error {
	// stop the trigger
	return nil
}