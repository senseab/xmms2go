package xmms2go

/*
We need import some enums from C.
Go should not use contents of native C directly
*/

const (
	IPCObjectSignal = iota
	IPCObjectMain
	IPCObjectPlaylist
	IPCObjectConfig
	IPCObjectPlayback
	IPCObjectMedialib
	IPCObjectCollection
	IPCObjectVisualization
	IPCObjectMediaInfoReader
	IPCObjectXform
	IPCObjectBindata
	IPCObjectCollSync
	IPCObjectCourier
	IPCObjectIPCManager
	IPCObjectEnd
)

const (
	IPCSignalPlaylistChanged = iota
	IPCSignalConfigValueChanged
	IPCSignalPlaybackStatus
	IPCSignalPlaybackVolumeChanged
	IPCSignalPlaybackPlaytime
	IPCSignalPlaybackCurrentID
	IPCSignalPlaylistCurrentPos
	IPCSignalPlaylistLoaded
	IPCSignalMedialibEntryAdded
	IPCSignalMedialibEntryUpdate
	IPCSignalMedialibEntryRemoved
	IPCSignalCollectionChanged
	IPCSignalQuit
	IPCSignalMediaInfoReaderStstus
	IPCSignalMediaInfoReaderUnindexed
	IPCSignalCourierMessage
	IPCSignalIPCManagerClientConnected
	IPCSignalIPCManagerClientDisconnected
	IPCSignalEnd
)

const IPCCmdFirst = 32

const (
	IPCCmdReply = iota
	IPCCmdError
)

/* Signal subsystem methods */
const (
	IPCCmdSignal = IPCCmdFirst + iota
	IPCCmdBroadcast
)

/* Main methods */
const (
	IPCCmdHello = IPCCmdFirst + iota
	IPCCmdQuit
	IPCCmdPluginList
	IPCCmdStats
)

/* Playlist methods */
const (
	IPCCmdREPLACE = IPCCmdFirst + iota
	IPCCmdSETPOS
	IPCCmdSETPOSREL
	IPCCmdADDURL
	IPCCmdADDCOLL
	IPCCmdREMOVEENTRY
	IPCCmdMOVEENTRY
	IPCCmdLIST
	IPCCmdCURRENTPOS
	IPCCmdCURRENTACTIVE
	IPCCmdINSERTURL
	IPCCmdINSERTCOLL
	IPCCmdLOAD
	IPCCmdRADD
	IPCCmdRINSERT
)

/* Config methods */
const (
	IPCCmdGETVALUE = IPCCmdFirst + iota
	IPCCmdSETVALUE
	IPCCmdREGVALUE
	IPCCmdLISTVALUES
)

/* playback methods */
const (
	IPCCmdSTART = IPCCmdFirst + iota
	IPCCmdSTOP
	IPCCmdPAUSE
	IPCCmdDECODERKILL
	IPCCmdCPLAYTIME
	IPCCmdSEEKMS
	IPCCmdSEEKSAMPLES
	IPCCmdPLAYBACKSTATUS
	IPCCmdCURRENTID
	IPCCmdVOLUMESET
	IPCCmdVOLUMEGET
)

/* Medialib methods */
const (
	IPCCmdINFO = IPCCmdFirst + iota
	IPCCmdPATHIMPORT
	IPCCmdREHASH
	IPCCmdGETID
	IPCCmdREMOVEID
	IPCCmdPROPERTYSETSTR
	IPCCmdPROPERTYSETINT
	IPCCmdPROPERTYREMOVE
	IPCCmdMOVEURL
	IPCCmdMLIBADDURL
)

/* Coll sync methods */
const (
	IPCCmdCOLLSYNCSYNC = IPCCmdFirst + iota
)

/* Collection methods */
const (
	IPCCmdCOLLECTIONGET = IPCCmdFirst + iota
	IPCCmdCOLLECTIONLIST
	IPCCmdCOLLECTIONSAVE
	IPCCmdCOLLECTIONREMOVE
	IPCCmdCOLLECTIONFIND
	IPCCmdCOLLECTIONRENAME
	IPCCmdQUERY
	IPCCmdQUERYINFOS
	IPCCmdIDLISTFROMPLS
)

/* bindata methods */
const (
	IPCCmdGETDATA = IPCCmdFirst + iota
	IPCCmdADDDATA
	IPCCmdREMOVEDATA
	IPCCmdLISTDATA
)

/* visualization methods */
const (
	IPCCmdVISUALIZATIONQUERYVERSION = IPCCmdFirst + iota
	IPCCmdVISUALIZATIONREGISTER
	IPCCmdVISUALIZATIONINITSHM
	IPCCmdVISUALIZATIONINITUDP
	IPCCmdVISUALIZATIONPROPERTY
	IPCCmdVISUALIZATIONPROPERTIES
	IPCCmdVISUALIZATIONSHUTDOWN
)

/* xform methods */
const (
	IPCCmdBROWSE = IPCCmdFirst + iota
)

/* courier methods */
const (
	IPCCmdSendMessage = IPCCmdFirst + iota
	IPCCmdReplyMessage
	IPCCmdGetConnectedClients
)

const (
	PlaylistChangedAdd = iota
	PlaylistChangedInsert
	PlaylistChangedShuffle /* deprecated */
	PlaylistChangedRemove
	PlaylistChangedClear /* deprecated */
	PlaylistChangedMove
	PlaylistChangedSort /* deprecated */
	PlaylistChangedUpdate
	PlaylistChangedReplace
)

const (
	CollectionChangedAdd = iota
	CollectionChangedUpdate
	CollectionChangedRename
	CollectionChangedRemove
)

const (
	PlaylistCurrentIDForget = iota
	PlaylistCurrentIDKeep
	PlaylistCurrentIDMoveToFront
)

const (
	PlaybackStatusStop = iota
	PlaybackStatusPlay
	PlaybackStatusPause
)

const (
	PlaybackSeekCur = 1 + iota
	PlaybackSeekSet
)

const (
	MediaInfoReaderStatusIdle = iota
	MediaInfoReaderStatusRunning
)

const (
	PluginTypeAll = iota
	PluginTypeOutput
	PluginTypeXform
)

/* Collection type */
const (
	CollectionTypeReference = iota
	CollectionTypeUniverse
	CollectionTypeUnion
	CollectionTypeIntersection
	CollectionTypeComplement
	CollectionTypeHas
	CollectionTypeMatch
	CollectionTypeToken
	CollectionTypeEquals
	CollectionTypeNotEqual
	CollectionTypeSmaller
	CollectionTypeSmallerEQ
	CollectionTypeGreater
	CollectionTypeGreaterEQ
	CollectionTypeOrder
	CollectionTypeLimit
	CollectionTypeMediaset
	CollectionTypeIDList
	CollectionTypeLast = CollectionTypeIDList
)

const (
	MedialibEntryStatusNew = iota
	MedialibEntryStatusOK
	MedialibEntryStatusResolving
	MedialibEntryStatusNotAvailable
	MedialibEntryStatusRehash
)

const (
	LogLevelUnknown = iota
	LogLevelFatal
	LogLevelFail
	LogLevelError
	LogLevelInfo
	LogLevelDebug
	LogLevelCount /* must be last */
)

const (
	C2CReplyPolicyNoReply = iota
	C2CReplyPolicySingleReply
	C2CReplyPolicyMultiReply
)
