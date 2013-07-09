Formatting library
==================

Formatting is a small (one file) library for formatting text. Currently it only
contains the IndentingWriter class, which is used to write text with variable
amounts of whitespace at the beginning of a line. Generally this would be used
to format source code, like this:

    buffer := bytes.NewBuffer(make([]byte, 0))
    ind := formatting.NewIndentingWriter(buffer)

    ind.Println("for {")
	ind.IncrIndent()
	ind.Println("// Do some stuff")
	ind.DecrIndent()
	ind.Println("}")


Copyright / License
===================
Copyright 2013 Travis Simon

Licensed under the GNU General Public License Version 2.0 (or later); you may not use this work except in compliance with the License. You may obtain a copy of the License in the LICENSE file, or at:

http://www.gnu.org/licenses/old-licenses/gpl-2.0.txt

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.