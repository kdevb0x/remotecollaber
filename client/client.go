package remotecollaber

import (
	"bufio"
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync/atomic"


	"github.com/gorilla/websocket"

	bp "github.com/faiface/beep"
	spk "github.com/faiface/beep/speaker"
	wav "github.com/faiface/beep/wav"
)

func connect(ctx url net.Listener, data *bp.Streamer) error {

}
