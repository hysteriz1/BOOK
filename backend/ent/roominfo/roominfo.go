// Code generated by entc, DO NOT EDIT.

package roominfo

const (
	// Label holds the string label denoting the roominfo type in the database.
	Label = "room_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldINFOBED holds the string denoting the infobed field in the database.
	FieldINFOBED = "infobed"
	// FieldINFOREFRIGERAT holds the string denoting the inforefrigerat field in the database.
	FieldINFOREFRIGERAT = "inforefrigerat"
	// FieldINFOWARDROB holds the string denoting the infowardrob field in the database.
	FieldINFOWARDROB = "infowardrob"
	// FieldINFOOFFICEDE holds the string denoting the infoofficede field in the database.
	FieldINFOOFFICEDE = "infoofficede"

	// EdgeINFOROOM holds the string denoting the info_room edge name in mutations.
	EdgeINFOROOM = "INFO_ROOM"

	// Table holds the table name of the roominfo in the database.
	Table = "room_infos"
	// INFOROOMTable is the table the holds the INFO_ROOM relation/edge.
	INFOROOMTable = "rooms"
	// INFOROOMInverseTable is the table name for the Room entity.
	// It exists in this package in order to avoid circular dependency with the "room" package.
	INFOROOMInverseTable = "rooms"
	// INFOROOMColumn is the table column denoting the INFO_ROOM relation/edge.
	INFOROOMColumn = "room_info_info_room"
)

// Columns holds all SQL columns for roominfo fields.
var Columns = []string{
	FieldID,
	FieldINFOBED,
	FieldINFOREFRIGERAT,
	FieldINFOWARDROB,
	FieldINFOOFFICEDE,
}
