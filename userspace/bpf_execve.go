// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//    ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗
//    ████╗  ██║██╔═████╗██║   ██║██╔══██╗
//    ██╔██╗ ██║██║██╔██║██║   ██║███████║
//    ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║
//    ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║
//    ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝

package userspace

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/cilium/ebpf/perf"
)

func EventExecve(event perf.Record) (*execve_data_t, error) {
	buffer := bytes.NewBuffer(event.RawSample)
	var data execve_data_t
	err := binary.Read(buffer, binary.LittleEndian, &data)
	if err != nil {
		return nil, fmt.Errorf("execve() kernel event perf: %v", err)
	}
	return &data, nil
}

type execve_data_t struct {
	Pid      uint32
	Filename [32]byte
	Comm     [32]byte
}
