package main

import (
	"fmt"

	"github.com/google/simhospital/pkg/hl7"
)

func main() {
	// Create a new HL7 message
	hl7Msg := `MSH|^~\&|SIMHOSP|SFAC|RAPP|RFAC|20200501140643||ORU^R01|1|T|2.3|||AL||44|ASCII|PID|1|2590157853^^^SIMULATOR MRN^MRN|2590157853^^^SIMULATOR MRN^MRN~2478684691^^^NHSNBR^NHSNMBR||Esterkin^AKI Scenario 6^^^Miss^^CURRENT||19890118000000|F|||170 Juice Place^^London^^RW21 6KC^GBR^HOME||020 5368 1665^HOME|||||||||R^Other - Chinese^^^|||||||||PV1|1|O|ED^^^Simulated Hospital^^ED^^|28b|||C006^Woolfson^Kathleen^^^Dr^^^DRNBR^PRSNL^^^ORGDR|||MED||||||||||||||||||||||||||||||||||20200501140643|||ORC|RE|1892929505|4262718364||CM||||20200501140643|OBR|1|1892929505|4262718364|us-0003^UREA AND ELECTROLYTES^WinPath^^||20200501140643|20200501140643|||||||20200501140643||||||||20200501140643|||F||1|OBX|1|NM|tt-0003-01^Creatinine^WinPath^^||98.00|UMOLL|49 - 92|H|||F|||20200501140643|||NTE|0||Task cow administration||NTE|1||Grapefruit garlic resale camera|`

	// mo := hl7.NewParseMessageOptions()
	// mo.TimezoneLoc = time.UTC

	if err := hl7.TimezoneAndLocation("UTC"); err != nil {
		fmt.Println("Cannot configure HL7 timezone and location")
	}
	m, err := hl7.ParseMessageV2([]byte(hl7Msg))
	if err != nil {
		fmt.Printf("Error parsing HL7 message: %v", err)
	}
	if _, ok := m.(*hl7.ORU_R01v2); !ok {
		fmt.Printf("ParseMessageV2 returned value of type %T, want *ORU_R01v2", m)
	}
	oru := m.(*hl7.ORU_R01v2)
	if oru.MSH() == nil {
		fmt.Printf("ParseMessageType.oru.MSH is <nil>, want not nil")
	}
	pids := oru.GroupByPID()
	if pids == nil {
		fmt.Printf("ParseMessageType.oru.GroupByPID is <nil>, want not nil")
	}
	// if got, want := len(pids), 1; got != want {
	// 	fmt.Printf("len(oru.GroupByPID) = %d, want %d", got, want)
	// }
	// if got, want := *pids[0].PID().Religion.Identifier, ST("CATHOLIC"); got != want {
	// 	fmt.Printf("PID.Religion got %q, want %q", got, want)
	// }
}
