package main

import rt "sky-app/rt"

var _ = rt.AsInt

func init() {
	rt.SetPortDefault("8000")
	rt.SetEnvDefault("SKY_LIVE_TTL", "1800")
	rt.SetEnvDefault("SKY_AUTH_TOKEN_TTL", "86400")
	rt.SetEnvDefault("SKY_AUTH_COOKIE", "sky_auth")
	rt.SetEnvDefault("SKY_AUTH_DRIVER", "jwt")
}

type Log_Config_Section = rt.SkyADT

var Log_Config_Section_NoSection Log_Config_Section = Log_Config_Section{Tag: 0, SkyName: "NoSection"}

func Log_Config_Section_InSource(v0 Log_Config_Source_R) Log_Config_Section {
	return Log_Config_Section{Tag: 1, SkyName: "InSource", Fields: []any{v0}}
}

func Log_Config_Section_InWebhook(v0 Log_Config_WebhookConfig_R) Log_Config_Section {
	return Log_Config_Section{Tag: 2, SkyName: "InWebhook", Fields: []any{v0}}
}

func init() { rt.RegisterAdtTag("NoSection", 0); rt.RegisterAdtTag("InSource", 1); rt.RegisterAdtTag("InWebhook", 2); }

type Log_Config_AppConfig_R struct { Sources []Log_Config_Source_R; Webhook Log_Config_WebhookConfig_R }

func init() { rt.RegisterGobType(Log_Config_AppConfig_R{}) }

func Log_Config_AppConfig(p0 []Log_Config_Source_R, p1 Log_Config_WebhookConfig_R) Log_Config_AppConfig_R { return Log_Config_AppConfig_R{Sources: p0, Webhook: p1} }

type Log_Config_ParseState_R struct { Sources []Log_Config_Source_R; Webhook Log_Config_WebhookConfig_R; Section Log_Config_Section }

func init() { rt.RegisterGobType(Log_Config_ParseState_R{}) }

func Log_Config_ParseState(p0 []Log_Config_Source_R, p1 Log_Config_WebhookConfig_R, p2 Log_Config_Section) Log_Config_ParseState_R { return Log_Config_ParseState_R{Sources: p0, Webhook: p1, Section: p2} }

type Log_Config_Source_R struct { Name string; Command string; Filter string; WebhookUrl string }

func init() { rt.RegisterGobType(Log_Config_Source_R{}) }

func Log_Config_Source(p0 string, p1 string, p2 string, p3 string) Log_Config_Source_R { return Log_Config_Source_R{Name: p0, Command: p1, Filter: p2, WebhookUrl: p3} }

type Log_Config_WebhookConfig_R struct { Url string; Filter string }

func init() { rt.RegisterGobType(Log_Config_WebhookConfig_R{}) }

func Log_Config_WebhookConfig(p0 string, p1 string) Log_Config_WebhookConfig_R { return Log_Config_WebhookConfig_R{Url: p0, Filter: p1} }

func Log_Config_emptySource() Log_Config_Source_R {
	return rt.Coerce[Log_Config_Source_R](Log_Config_Source_R{Command: any("").(string), Filter: any("").(string), Name: any("").(string), WebhookUrl: any("").(string)})
}

func Log_Config_emptyWebhook() Log_Config_WebhookConfig_R {
	return rt.Coerce[Log_Config_WebhookConfig_R](Log_Config_WebhookConfig_R{Filter: any("").(string), Url: any("").(string)})
}

func Log_Config_parseConfig(path string) Log_Config_AppConfig_R {
	return rt.Coerce[Log_Config_AppConfig_R](func() any { __subject := rt.ResultCoerce[any, any](rt.AnyTaskRun(rt.File_readFile(path))); if __subject.Tag == 0 {; 	content := rt.ResultOk(any(__subject)); 	_ = content; 	return Log_Config_parseLines(any(Log_Config_ParseState_R{Section: any(Log_Config_Section_NoSection).(Log_Config_Section), Sources: rt.AsListT[Log_Config_Source_R]([]any{}), Webhook: any(Log_Config_emptyWebhook()).(Log_Config_WebhookConfig_R)}).(Log_Config_ParseState_R), rt.AsListT[string](rt.String_lines(content))); }; if __subject.Tag == 1 {; 	_ = rt.ResultErr(any(__subject)); 	return Log_Config_AppConfig_R{Sources: rt.AsListT[Log_Config_Source_R]([]any{}), Webhook: any(Log_Config_emptyWebhook()).(Log_Config_WebhookConfig_R)}; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Config_finalizeState(state Log_Config_ParseState_R) Log_Config_AppConfig_R {
	return rt.Coerce[Log_Config_AppConfig_R](func() any { __subject := any(rt.Field(state, "Section")).(Log_Config_Section); if __subject.Tag == 1 {; 	s := rt.AdtField(any(__subject), 0); 	_ = s; 	return Log_Config_AppConfig_R{Sources: rt.AsListT[Log_Config_Source_R](rt.List_append(rt.Field(state, "Sources"), []any{s})), Webhook: any(rt.Field(state, "Webhook")).(Log_Config_WebhookConfig_R)}; }; if __subject.Tag == 2 {; 	w := rt.AdtField(any(__subject), 0); 	_ = w; 	return Log_Config_AppConfig_R{Sources: rt.AsListT[Log_Config_Source_R](rt.Field(state, "Sources")), Webhook: any(w).(Log_Config_WebhookConfig_R)}; }; if __subject.Tag == 0 {; 	return Log_Config_AppConfig_R{Sources: rt.AsListT[Log_Config_Source_R](rt.Field(state, "Sources")), Webhook: any(rt.Field(state, "Webhook")).(Log_Config_WebhookConfig_R)}; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Config_parseLines(state Log_Config_ParseState_R, lines []string) Log_Config_AppConfig_R {
	return rt.Coerce[Log_Config_AppConfig_R](func() any { __subject := lines; if len(rt.AsList(__subject)) == 0 {; 	return Log_Config_finalizeState(any(state).(Log_Config_ParseState_R)); }; if len(rt.AsList(__subject)) >= 1 {; 	line := rt.AsList(__subject)[0]; 	_ = line; 	rest := any(rt.AsList(__subject)[1:]); 	_ = rest; 	return func() any { trimmed := rt.String_trimT(rt.AsString(line)); _ = trimmed; return func() any { if rt.AsBool(rt.Eq(trimmed, "[[source]]")) {; 	return func() any { newState := func() any { __subject := any(rt.Field(state, "Section")).(Log_Config_Section); if __subject.Tag == 1 {; 	s := rt.AdtField(any(__subject), 0); 	_ = s; 	return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InSource(any(Log_Config_emptySource()).(Log_Config_Source_R)), "Sources": rt.List_append(rt.Field(state, "Sources"), []any{s})}); }; if __subject.Tag == 2 {; 	w := rt.AdtField(any(__subject), 0); 	_ = w; 	return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InSource(any(Log_Config_emptySource()).(Log_Config_Source_R)), "Webhook": w}); }; if __subject.Tag == 0 {; 	return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InSource(any(Log_Config_emptySource()).(Log_Config_Source_R))}); }; _ = rt.Unreachable("case/__subject"); return nil }(); _ = newState; return Log_Config_parseLines(any(newState).(Log_Config_ParseState_R), rt.AsListT[string](rest)) }(); } else {; 	if rt.AsBool(rt.Eq(trimmed, "[webhook]")) {; 		return func() any { newState := func() any { __subject := any(rt.Field(state, "Section")).(Log_Config_Section); if __subject.Tag == 1 {; 	s := rt.AdtField(any(__subject), 0); 	_ = s; 	return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InWebhook(any(rt.Field(state, "Webhook")).(Log_Config_WebhookConfig_R)), "Sources": rt.List_append(rt.Field(state, "Sources"), []any{s})}); }; return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InWebhook(any(rt.Field(state, "Webhook")).(Log_Config_WebhookConfig_R))}); _ = rt.Unreachable("case/__subject"); return nil }(); _ = newState; return Log_Config_parseLines(any(newState).(Log_Config_ParseState_R), rt.AsListT[string](rest)) }(); 	} else {; 		return func() any { newState := func() any { __subject := any(rt.Field(state, "Section")).(Log_Config_Section); if __subject.Tag == 1 {; 	s := rt.AdtField(any(__subject), 0); 	_ = s; 	return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InSource(any(Log_Config_parseSourceLine(rt.CoerceString(trimmed), any(s).(Log_Config_Source_R))).(Log_Config_Source_R))}); }; if __subject.Tag == 2 {; 	w := rt.AdtField(any(__subject), 0); 	_ = w; 	return rt.RecordUpdate(state, map[string]any{"Section": Log_Config_Section_InWebhook(any(Log_Config_parseWebhookLine(rt.CoerceString(trimmed), any(w).(Log_Config_WebhookConfig_R))).(Log_Config_WebhookConfig_R))}); }; if __subject.Tag == 0 {; 	return state; }; _ = rt.Unreachable("case/__subject"); return nil }(); _ = newState; return Log_Config_parseLines(any(newState).(Log_Config_ParseState_R), rt.AsListT[string](rest)) }(); 	}; }; return nil }() }(); }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Config_parseSourceLine(line string, source Log_Config_Source_R) Log_Config_Source_R {
	return rt.Coerce[Log_Config_Source_R](func() any { if rt.AsBool(rt.String_startsWithT("name", rt.AsString(line))) {; 	return rt.RecordUpdate(source, map[string]any{"Name": Log_Config_extractValue(rt.CoerceString(line))}); } else {; 	if rt.AsBool(rt.String_startsWithT("command", rt.AsString(line))) {; 		return rt.RecordUpdate(source, map[string]any{"Command": Log_Config_extractValue(rt.CoerceString(line))}); 	} else {; 		if rt.AsBool(rt.String_startsWithT("filter", rt.AsString(line))) {; 			return rt.RecordUpdate(source, map[string]any{"Filter": Log_Config_extractValue(rt.CoerceString(line))}); 		} else {; 			if rt.AsBool(rt.String_startsWithT("webhook_url", rt.AsString(line))) {; 				return rt.RecordUpdate(source, map[string]any{"WebhookUrl": Log_Config_extractValue(rt.CoerceString(line))}); 			} else {; 				return source; 			}; 		}; 	}; }; return nil }())
}

func Log_Config_parseWebhookLine(line string, webhook Log_Config_WebhookConfig_R) Log_Config_WebhookConfig_R {
	return rt.Coerce[Log_Config_WebhookConfig_R](func() any { if rt.AsBool(rt.String_startsWithT("url", rt.AsString(line))) {; 	return rt.RecordUpdate(webhook, map[string]any{"Url": Log_Config_extractValue(rt.CoerceString(line))}); } else {; 	if rt.AsBool(rt.String_startsWithT("filter", rt.AsString(line))) {; 		return rt.RecordUpdate(webhook, map[string]any{"Filter": Log_Config_extractValue(rt.CoerceString(line))}); 	} else {; 		return webhook; 	}; }; return nil }())
}

func Log_Config_extractValue(line string) string {
	return rt.CoerceString(func() any { __subject := rt.String_splitT("=", rt.AsString(line)); if len(rt.AsList(__subject)) == 0 {; 	return ""; }; if len(rt.AsList(__subject)) >= 1 {; 	_ = rt.AsList(__subject)[0]; 	rest := any(rt.AsList(__subject)[1:]); 	_ = rest; 	return Log_Config_stripQuotes(rt.CoerceString(rt.String_trimT(rt.AsString(rt.String_join("=", rest))))); }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Config_stripQuotes(s string) string {
	return rt.CoerceString(func() any { len := rt.String_lengthT(rt.AsString(s)); _ = len; return func() any { if rt.AsBool(rt.And(rt.String_startsWithT("\"", rt.AsString(s)), rt.String_endsWithT("\"", rt.AsString(s)))) {; 	return rt.String_sliceT(1, rt.AsInt(rt.Sub(len, 1)), rt.AsString(s)); } else {; 	return s; }; return nil }() }())
}

type Log_Entry_Level = int

const (
	Log_Entry_Level_Debug Log_Entry_Level = iota
	Log_Entry_Level_Info
	Log_Entry_Level_Warn
	Log_Entry_Level_ErrorLevel
)

type Log_Entry_LogEntry_R struct { Timestamp int; Level Log_Entry_Level; Scope string; Message string; Source string }

func init() { rt.RegisterGobType(Log_Entry_LogEntry_R{}) }

func Log_Entry_LogEntry(p0 int, p1 Log_Entry_Level, p2 string, p3 string, p4 string) Log_Entry_LogEntry_R { return Log_Entry_LogEntry_R{Timestamp: p0, Level: p1, Scope: p2, Message: p3, Source: p4} }

func Log_Entry_levelFromString(str string) Log_Entry_Level {
	return rt.Coerce[Log_Entry_Level](func() any { __subject := rt.String_toLowerT(rt.AsString(str)); if __subject == "debug" {; 	return Log_Entry_Level_Debug; }; if __subject == "info" {; 	return Log_Entry_Level_Info; }; if __subject == "warn" {; 	return Log_Entry_Level_Warn; }; if __subject == "error" {; 	return Log_Entry_Level_ErrorLevel; }; return Log_Entry_Level_Info; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Entry_levelToString(level Log_Entry_Level) string {
	return rt.CoerceString(func() any { __subject := level; if rt.EnumTagIs(__subject, 0) {; 	return "DEBUG"; }; if rt.EnumTagIs(__subject, 1) {; 	return "INFO"; }; if rt.EnumTagIs(__subject, 2) {; 	return "WARN"; }; if rt.EnumTagIs(__subject, 3) {; 	return "ERROR"; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Entry_levelToInt(level Log_Entry_Level) int {
	return rt.CoerceInt(func() any { __subject := level; if rt.EnumTagIs(__subject, 0) {; 	return 0; }; if rt.EnumTagIs(__subject, 1) {; 	return 1; }; if rt.EnumTagIs(__subject, 2) {; 	return 2; }; if rt.EnumTagIs(__subject, 3) {; 	return 3; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Entry_levelDecoder() rt.SkyDecoder {
	return rt.Coerce[rt.SkyDecoder](rt.JsonDec_map(Log_Entry_levelFromString, rt.JsonDec_string()))
}

func Log_Entry_entryDecoder() rt.SkyDecoder {
	return rt.Coerce[rt.SkyDecoder](rt.JsonDecP_required("message", rt.JsonDec_string(), rt.JsonDecP_required("scope", rt.JsonDec_string(), rt.JsonDecP_required("level", Log_Entry_levelDecoder(), rt.JsonDecP_required("timestamp", rt.JsonDec_int(), rt.JsonDec_succeed(func(t any) any { return func(l any) any { return func(s any) any { return func(m any) any { return Log_Entry_LogEntry_R{Level: any(l).(Log_Entry_Level), Message: any(m).(string), Scope: any(s).(string), Source: any("").(string), Timestamp: any(t).(int)}; }; }; }; }))))))
}

func Log_Entry_decodeEntry(line string, source string) Log_Entry_LogEntry_R {
	return rt.Coerce[Log_Entry_LogEntry_R](func() any { __subject := rt.ResultCoerce[any, any](rt.JsonDec_decodeString(Log_Entry_entryDecoder(), line)); if __subject.Tag == 0 {; 	entry := rt.ResultOk(any(__subject)); 	_ = entry; 	return rt.RecordUpdate(entry, map[string]any{"Source": source}); }; if __subject.Tag == 1 {; 	_ = rt.ResultErr(any(__subject)); 	return Log_Entry_parseRawLine(rt.CoerceString(line), rt.CoerceString(source)); }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func Log_Entry_parseRawLine(line string, source string) Log_Entry_LogEntry_R {
	return rt.Coerce[Log_Entry_LogEntry_R](func() any { ts := Log_Entry_extractTimestamp(rt.CoerceString(line)); _ = ts; return func() any { level := Log_Entry_detectLevel(rt.CoerceString(line)); _ = level; return func() any { scope := func() any { if rt.AsBool(rt.String_containsT("|", rt.AsString(line))) {; 	return rt.String_trimT(rt.AsString(func() any { __subject := rt.String_splitT("|", rt.AsString(line)); if len(rt.AsList(__subject)) >= 1 {; 	_ = rt.AsList(__subject)[0]; 	rest := any(rt.AsList(__subject)[1:]); 	_ = rest; 	return ""; }; return ""; _ = rt.Unreachable("case/__subject"); return nil }())); } else {; 	return "raw"; }; return nil }(); _ = scope; return Log_Entry_LogEntry_R{Level: any(level).(Log_Entry_Level), Message: any(line).(string), Scope: any(source).(string), Source: any(source).(string), Timestamp: any(ts).(int)} }() }() }())
}

func Log_Entry_extractTimestamp(line string) int {
	return rt.CoerceInt(func() any { prefix := rt.String_sliceT(0, 19, rt.AsString(line)); _ = prefix; return func() any { hasIsoDate := rt.And(rt.Gte(rt.String_lengthT(rt.AsString(prefix)), 19), rt.And(rt.String_containsT("-", rt.AsString(rt.String_sliceT(0, 10, rt.AsString(prefix)))), rt.String_containsT(":", rt.AsString(rt.String_sliceT(11, 19, rt.AsString(prefix)))))); _ = hasIsoDate; return func() any { if rt.AsBool(hasIsoDate) {; 	return func() any { timePart := rt.String_sliceT(11, 19, rt.AsString(prefix)); _ = timePart; return func() any { parts := rt.String_splitT(":", rt.AsString(timePart)); _ = parts; return func() any { h := rt.Maybe_withDefaultAnyT(0, any(rt.String_toInt(rt.Maybe_withDefaultAnyT("0", any(rt.List_headAny(any(parts))))))); _ = h; return func() any { m := rt.Maybe_withDefaultAnyT(0, any(rt.String_toInt(rt.Maybe_withDefaultAnyT("0", any(rt.List_headAny(any(rt.List_dropAnyT(1, rt.AsList(parts))))))))); _ = m; return func() any { s := rt.Maybe_withDefaultAnyT(0, any(rt.String_toInt(rt.Maybe_withDefaultAnyT("0", any(rt.List_headAny(any(rt.List_dropAnyT(2, rt.AsList(parts))))))))); _ = s; return rt.Add(rt.Add(rt.Mul(h, 3600), rt.Mul(m, 60)), s) }() }() }() }() }(); } else {; 	return 0; }; return nil }() }() }())
}

func Log_Entry_detectLevel(line string) Log_Entry_Level {
	return rt.Coerce[Log_Entry_Level](func() any { lower := rt.String_toLowerT(rt.AsString(line)); _ = lower; return func() any { if rt.AsBool(rt.Or(rt.String_containsT("error", rt.AsString(lower)), rt.String_containsT("err", rt.AsString(lower)))) {; 	return Log_Entry_Level_ErrorLevel; } else {; 	if rt.AsBool(rt.String_containsT("warn", rt.AsString(lower))) {; 		return Log_Entry_Level_Warn; 	} else {; 		if rt.AsBool(rt.String_containsT("debug", rt.AsString(lower))) {; 			return Log_Entry_Level_Debug; 		} else {; 			return Log_Entry_Level_Info; 		}; 	}; }; return nil }() }())
}

type Log_Webhook_WebhookRule_R struct { Filter string; Url string; SourceName string }

func init() { rt.RegisterGobType(Log_Webhook_WebhookRule_R{}) }

func Log_Webhook_WebhookRule(p0 string, p1 string, p2 string) Log_Webhook_WebhookRule_R { return Log_Webhook_WebhookRule_R{Filter: p0, Url: p1, SourceName: p2} }

func Log_Webhook_logMsg(msg string) struct{} {
	return rt.Coerce[struct{}](func() any { _ = rt.AnyTaskRun(rt.Go_Os_fileWriteStringT(rt.Coerce[rt.FfiT_Go_Os_fileWriteString_P0](rt.Go_Os_stderr(struct{}{})), rt.CoerceString(rt.Concat(msg, "\n")))); return struct{}{} }())
}

func Log_Webhook_buildRules(config Log_Config_AppConfig_R) []Log_Webhook_WebhookRule_R {
	return rt.Coerce[[]Log_Webhook_WebhookRule_R](rt.List_filterMapAnyT(any(func(source any) any { return func() any { filter := func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(rt.Field(source, "Filter")))) {; 	return rt.Field(rt.Field(config, "Webhook"), "Filter"); } else {; 	return rt.Field(source, "Filter"); }; return nil }(); _ = filter; return func() any { url := func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(rt.Field(source, "WebhookUrl")))) {; 	return rt.Field(rt.Field(config, "Webhook"), "Url"); } else {; 	return rt.Field(source, "WebhookUrl"); }; return nil }(); _ = url; return func() any { if rt.AsBool(rt.Or(rt.String_isEmptyT(rt.AsString(filter)), rt.String_isEmptyT(rt.AsString(url)))) {; 	return rt.Nothing[any](); } else {; 	return rt.Just[any](Log_Webhook_WebhookRule_R{Filter: any(filter).(string), SourceName: any(rt.Field(source, "Name")).(string), Url: any(url).(string)}); }; return nil }() }() }(); }), rt.AsList(rt.Field(config, "Sources"))))
}

func Log_Webhook_matchesFilter(pattern string, entry Log_Entry_LogEntry_R) bool {
	return rt.CoerceBool(func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(pattern))) {; 	return false; } else {; 	return func() any { __subject_tFfi := rt.Go_Regexp_compileT(rt.CoerceString(rt.Concat("(?i)", pattern))); if __subject_tFfi.Tag == 0 {; 	re := any(__subject_tFfi.OkValue); 	_ = re; 	return rt.Result_withDefaultAnyT(any(false), any(rt.Go_Regexp_regexpMatchStringT(rt.Coerce[rt.FfiT_Go_Regexp_regexpMatchString_P0](re), rt.CoerceString(rt.Field(entry, "Message"))))); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.String_containsT(rt.AsString(rt.String_toLowerT(rt.AsString(pattern))), rt.AsString(rt.String_toLowerT(rt.AsString(rt.Field(entry, "Message"))))); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }(); }; return nil }())
}

func Log_Webhook_buildPayload(entry Log_Entry_LogEntry_R) string {
	return rt.CoerceString(rt.JsonEnc_encode(0, rt.JsonEnc_object([]any{rt.SkyTuple2{V0: "source", V1: rt.JsonEnc_string(rt.Field(entry, "Source"))}, rt.SkyTuple2{V0: "level", V1: rt.JsonEnc_string(Log_Entry_levelToString(any(rt.Field(entry, "Level")).(Log_Entry_Level)))}, rt.SkyTuple2{V0: "scope", V1: rt.JsonEnc_string(rt.Field(entry, "Scope"))}, rt.SkyTuple2{V0: "message", V1: rt.JsonEnc_string(rt.Field(entry, "Message"))}, rt.SkyTuple2{V0: "timestamp", V1: rt.JsonEnc_int(rt.Field(entry, "Timestamp"))}})))
}

func Log_Webhook_send(url string, entry Log_Entry_LogEntry_R) struct{} {
	return rt.Coerce[struct{}](func() any { payload := Log_Webhook_buildPayload(any(entry).(Log_Entry_LogEntry_R)); _ = payload; return func() any { cmd := rt.Go_Exec_commandT("curl", rt.Coerce[[]string]([]any{"-s", "-X", "POST", "-H", "Content-Type: application/json", "-d", payload, url})); _ = cmd; return func() any { _ = rt.AnyTaskRun(rt.Go_Exec_cmdRunT(rt.Coerce[rt.FfiT_Go_Exec_cmdRun_P0](cmd))); return struct{}{} }() }() }())
}

func Log_Webhook_processNewEntries(rules []Log_Webhook_WebhookRule_R, entries []Log_Entry_LogEntry_R) struct{} {
	return rt.Coerce[struct{}](func() any { if rt.AsBool(rt.List_isEmptyT(rt.AsList(rules))) {; 	return struct{}{}; } else {; 	return rt.List_foldlAnyT(any(func(entry any) any { return func(_ any) any { return rt.List_foldlAnyT(any(func(rule any) any { return func(_ any) any { return func() any { if rt.AsBool(rt.And(rt.Eq(rt.Field(rule, "SourceName"), rt.Field(entry, "Source")), Log_Webhook_matchesFilter(rt.CoerceString(rt.Field(rule, "Filter")), any(entry).(Log_Entry_LogEntry_R)))) {; 	return func() any { _ = rt.AnyTaskRun(Log_Webhook_logMsg(rt.CoerceString(rt.Concat("[webhook] filter matched: source=", rt.Concat(rt.Field(entry, "Source"), rt.Concat(" filter=", rt.Concat(rt.Field(rule, "Filter"), rt.Concat(" message=", rt.Field(entry, "Message"))))))))); return func() any { _ = rt.AnyTaskRun(Log_Webhook_send(rt.CoerceString(rt.Field(rule, "Url")), any(entry).(Log_Entry_LogEntry_R))); return func() any { _ = rt.AnyTaskRun(Log_Webhook_logMsg(rt.CoerceString(rt.Concat("[webhook] fired: url=", rt.Field(rule, "Url"))))); return struct{}{} }() }() }(); } else {; 	return struct{}{}; }; return nil }(); }; }), any(struct{}{}), rt.AsList(rules)); }; }), any(struct{}{}), rt.AsList(entries)); }; return nil }())
}

type Msg = rt.SkyADT

var Msg_Tick Msg = Msg{Tag: 0, SkyName: "Tick"}

func Msg_SetLevelFilter(v0 string) Msg {
	return Msg{Tag: 1, SkyName: "SetLevelFilter", Fields: []any{v0}}
}

func Msg_SetScopeFilter(v0 string) Msg {
	return Msg{Tag: 2, SkyName: "SetScopeFilter", Fields: []any{v0}}
}

func Msg_SetSearchFilter(v0 string) Msg {
	return Msg{Tag: 3, SkyName: "SetSearchFilter", Fields: []any{v0}}
}

func Msg_SetSourceFilter(v0 string) Msg {
	return Msg{Tag: 4, SkyName: "SetSourceFilter", Fields: []any{v0}}
}

var Msg_ToggleAutoScroll Msg = Msg{Tag: 5, SkyName: "ToggleAutoScroll"}

var Msg_ToggleTheme Msg = Msg{Tag: 6, SkyName: "ToggleTheme"}

var Msg_ClearLogs Msg = Msg{Tag: 7, SkyName: "ClearLogs"}

func Msg_Navigate(v0 Page) Msg {
	return Msg{Tag: 8, SkyName: "Navigate", Fields: []any{v0}}
}

func init() { rt.RegisterAdtTag("Tick", 0); rt.RegisterAdtTag("SetLevelFilter", 1); rt.RegisterAdtTag("SetScopeFilter", 2); rt.RegisterAdtTag("SetSearchFilter", 3); rt.RegisterAdtTag("SetSourceFilter", 4); rt.RegisterAdtTag("ToggleAutoScroll", 5); rt.RegisterAdtTag("ToggleTheme", 6); rt.RegisterAdtTag("ClearLogs", 7); rt.RegisterAdtTag("Navigate", 8); }

type Page = int

const (
	Page_LogPage Page = iota
)

type SourceMode = int

const (
	SourceMode_FileMode SourceMode = iota
	SourceMode_StreamMode
)

type Model_R struct {
	Entries []Log_Entry_LogEntry_R
	LevelFilter string
	ScopeFilter string
	SearchFilter string
	SourceFilter string
	AutoScroll bool
	Theme string
	Watched []WatchedFile_R
	Scanners []SourceScanner_R
	FileCounts map[string]int
	SourceMode SourceMode
	WebhookRules []Log_Webhook_WebhookRule_R
}

func init() { rt.RegisterGobType(Model_R{}) }

func Model(p0 []Log_Entry_LogEntry_R, p1 string, p2 string, p3 string, p4 string, p5 bool, p6 string, p7 []WatchedFile_R, p8 []SourceScanner_R, p9 map[string]int, p10 SourceMode, p11 []Log_Webhook_WebhookRule_R) Model_R { return Model_R{Entries: p0, LevelFilter: p1, ScopeFilter: p2, SearchFilter: p3, SourceFilter: p4, AutoScroll: p5, Theme: p6, Watched: p7, Scanners: p8, FileCounts: p9, SourceMode: p10, WebhookRules: p11} }

type SourceScanner_R struct {
	Scanner any
	Label string
}

func init() { rt.RegisterGobType(SourceScanner_R{}) }

func SourceScanner(p0 any, p1 string) SourceScanner_R { return SourceScanner_R{Scanner: p0, Label: p1} }

type WatchedFile_R struct {
	Path string
	Label string
}

func init() { rt.RegisterGobType(WatchedFile_R{}) }

func WatchedFile(p0 string, p1 string) WatchedFile_R { return WatchedFile_R{Path: p0, Label: p1} }

func helpText() string {
	return rt.CoerceString("sky-log — real-time log viewer\n\nUsage:\n  sky-log [options] [file ...]\n  command | sky-log\n\nOptions:\n  -h, --help    Show this help message\n\nExamples:\n  sky-log app.log                  Watch a log file\n  sky-log app.log server.log       Watch multiple files\n  tail -f /var/log/syslog | sky-log  Pipe stdin\n  sky-log                          Use sources from sky-log.toml\n")
}

func readFileEntries(source WatchedFile_R) []Log_Entry_LogEntry_R {
	return rt.Coerce[[]Log_Entry_LogEntry_R](func() any { __subject := rt.ResultCoerce[any, any](rt.AnyTaskRun(rt.File_readFile(rt.Field(source, "Path")))); if __subject.Tag == 0 {; 	content := rt.ResultOk(any(__subject)); 	_ = content; 	return rt.List_mapAny(any(func(l any) any { return Log_Entry_decodeEntry(rt.CoerceString(l), rt.CoerceString(rt.Field(source, "Label"))); }), any(rt.List_filterAny(any(func(l any) any { return rt.Basics_notT(rt.AsBool(rt.String_isEmptyT(rt.AsString(l)))); }), any(rt.String_lines(content))))); }; if __subject.Tag == 1 {; 	_ = rt.ResultErr(any(__subject)); 	return []any{}; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func collectFileEntries(watched []WatchedFile_R, counts map[string]int, existing []Log_Entry_LogEntry_R) any {
	return rt.List_foldlAnyT(any(func(source any) any { return func(acc any) any { return func() any { seen := func() any { __subject := rt.MaybeCoerce[any](rt.Dict_getAnyT(any(rt.Field(source, "Path")), any(rt.Field(acc, "Counts")))); if __subject.Tag == 0 {; 	n := rt.MaybeJust(any(__subject)); 	_ = n; 	return n; }; if __subject.Tag == 1 {; 	return 0; }; _ = rt.Unreachable("case/__subject"); return nil }(); _ = seen; return func() any { allEntries := readFileEntries(any(source).(WatchedFile_R)); _ = allEntries; return func() any { total := rt.List_lengthT(rt.AsList(allEntries)); _ = total; return func() any { newOnes := rt.List_dropAnyT(rt.AsInt(seen), rt.AsList(allEntries)); _ = newOnes; return struct{ Counts any; Entries any; NewEntries any }{Counts: rt.Dict_insertT(rt.AsString(rt.Field(source, "Path")), any(total), rt.AsDict(rt.Field(acc, "Counts"))), Entries: rt.List_append(rt.Field(acc, "Entries"), newOnes), NewEntries: rt.List_append(rt.Field(acc, "NewEntries"), newOnes)} }() }() }() }(); }; }), any(struct{ Counts any; Entries any; NewEntries any }{Counts: counts, Entries: existing, NewEntries: []any{}}), rt.AsList(watched))
}

func scanLine(source SourceScanner_R) rt.SkyMaybe[Log_Entry_LogEntry_R] {
	return rt.MaybeCoerce[Log_Entry_LogEntry_R](func() any { if rt.AsBool(rt.Result_withDefaultAnyT(any(false), any(rt.Go_Bufio_scannerScanT(rt.Coerce[rt.FfiT_Go_Bufio_scannerScan_P0](rt.Field(source, "Scanner")))))) {; 	return func() any { text := rt.Result_withDefaultAnyT("", any(rt.Go_Bufio_scannerTextT(rt.Coerce[rt.FfiT_Go_Bufio_scannerText_P0](rt.Field(source, "Scanner"))))); _ = text; return rt.Just[any](Log_Entry_decodeEntry(rt.CoerceString(text), rt.CoerceString(rt.Field(source, "Label")))) }(); } else {; 	return rt.Nothing[any](); }; return nil }())
}

func collectStreamEntries(scanners []SourceScanner_R) []Log_Entry_LogEntry_R {
	return rt.Coerce[[]Log_Entry_LogEntry_R](rt.List_filterMapAnyT(any(scanLine), rt.AsList(scanners)))
}

func initStdinScanner() rt.SkyMaybe[SourceScanner_R] {
	return rt.MaybeCoerce[SourceScanner_R](func() any { __subject_tFfi := rt.Go_Bufio_newScannerT(rt.Coerce[rt.FfiT_Go_Bufio_newScanner_P0](rt.Basics_identityT(any(rt.Go_Os_stdin(struct{}{}))))); if __subject_tFfi.Tag == 0 {; 	scanner := any(__subject_tFfi.OkValue); 	_ = scanner; 	return rt.Just[any](SourceScanner_R{Label: any("stdin").(string), Scanner: scanner}); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.Nothing[any](); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }())
}

func initCommandScanner(source Log_Config_Source_R) rt.SkyMaybe[SourceScanner_R] {
	return rt.MaybeCoerce[SourceScanner_R](func() any { __subject_tFfi := rt.Go_Exec_commandT("sh", rt.Coerce[[]string]([]any{"-c", rt.Field(source, "Command")})); if __subject_tFfi.Tag == 0 {; 	cmd := any(__subject_tFfi.OkValue); 	_ = cmd; 	return func() any { __subject_tFfi := rt.Go_Exec_cmdStdoutPipeT(rt.Coerce[rt.FfiT_Go_Exec_cmdStdoutPipe_P0](cmd)); if __subject_tFfi.Tag == 0 {; 	pipe := any(__subject_tFfi.OkValue); 	_ = pipe; 	return func() any { __subject_tFfi := rt.Go_Exec_cmdStartT(rt.Coerce[rt.FfiT_Go_Exec_cmdStart_P0](cmd)); if __subject_tFfi.Tag == 0 {; 	_ = any(__subject_tFfi.OkValue); 	return func() any { __subject_tFfi := rt.Go_Bufio_newScannerT(rt.Coerce[rt.FfiT_Go_Bufio_newScanner_P0](rt.Basics_identityT(any(pipe)))); if __subject_tFfi.Tag == 0 {; 	scanner := any(__subject_tFfi.OkValue); 	_ = scanner; 	return rt.Just[any](SourceScanner_R{Label: any(rt.Field(source, "Name")).(string), Scanner: scanner}); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.Nothing[any](); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }(); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.Nothing[any](); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }(); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.Nothing[any](); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }(); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.Nothing[any](); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }())
}

func resolveMode(args []string, sources []Log_Config_Source_R) any {
	return func() any { if rt.AsBool(rt.Basics_notT(rt.AsBool(rt.List_isEmptyT(rt.AsList(args))))) {; 	return struct{ Mode any; Scanners any; Watched any }{Mode: SourceMode_FileMode, Scanners: []any{}, Watched: rt.List_mapAny(any(func(f any) any { return WatchedFile_R{Label: any(f).(string), Path: any(f).(string)}; }), any(args))}; } else {; 	return func() any { cmdScanners := rt.List_filterMapAnyT(any(initCommandScanner), rt.AsList(sources)); _ = cmdScanners; return func() any { if rt.AsBool(rt.Basics_notT(rt.AsBool(rt.List_isEmptyT(rt.AsList(cmdScanners))))) {; 	return struct{ Mode any; Scanners any; Watched any }{Mode: SourceMode_StreamMode, Scanners: cmdScanners, Watched: []any{}}; } else {; 	return struct{ Mode any; Scanners any; Watched any }{Mode: SourceMode_StreamMode, Scanners: rt.List_filterMapAnyT(any(rt.Basics_identity[any]), rt.AsList([]any{initStdinScanner()})), Watched: []any{}}; }; return nil }() }(); }; return nil }()
}

func init_[T1 any](_ T1) rt.SkyTuple2 {
	return rt.Coerce[rt.SkyTuple2](func() any { rawArgs := rt.List_dropAnyT(1, rt.AsList(rt.Result_withDefaultAnyT(any([]any{}), any(rt.AnyTaskRun(rt.System_args(struct{}{})))))); _ = rawArgs; return func() any { _ = rt.AnyTaskRun(func() any { if rt.AsBool(rt.List_anyAnyT(any(func(a any) any { return rt.Or(rt.Eq(a, "--help"), rt.Eq(a, "-h")); }), rt.AsList(rawArgs))) {; 	return func() any { _ = rt.AnyTaskRun(rt.Log_printlnT(any(helpText()))); return rt.System_exit(0) }(); } else {; 	return struct{}{}; }; return nil }()); return func() any { args := rt.List_filterAny(any(func(a any) any { return rt.Basics_notT(rt.AsBool(rt.String_startsWithT("-", rt.AsString(a)))); }), any(rawArgs)); _ = args; return func() any { config := Log_Config_parseConfig(rt.CoerceString("sky-log.toml")); _ = config; return func() any { resolved := resolveMode(rt.AsListT[string](args), rt.AsListT[Log_Config_Source_R](rt.Field(config, "Sources"))); _ = resolved; return func() any { webhookRules := Log_Webhook_buildRules(any(config).(Log_Config_AppConfig_R)); _ = webhookRules; return func() any { result := collectFileEntries(rt.AsListT[WatchedFile_R](rt.Field(resolved, "Watched")), rt.AsMapT[int](rt.Dict_empty()), rt.AsListT[Log_Entry_LogEntry_R]([]any{})); _ = result; return rt.SkyTuple2{V0: Model_R{AutoScroll: any(true).(bool), Entries: rt.AsListT[Log_Entry_LogEntry_R](rt.Field(result, "Entries")), FileCounts: rt.AsMapT[int](rt.Field(result, "Counts")), LevelFilter: any("").(string), Scanners: rt.AsListT[SourceScanner_R](rt.Field(resolved, "Scanners")), ScopeFilter: any("").(string), SearchFilter: any("").(string), SourceFilter: any("").(string), SourceMode: any(rt.Field(resolved, "Mode")).(SourceMode), Theme: any("dark").(string), Watched: rt.AsListT[WatchedFile_R](rt.Field(resolved, "Watched")), WebhookRules: rt.AsListT[Log_Webhook_WebhookRule_R](webhookRules)}, V1: rt.Cmd_none()} }() }() }() }() }() }() }())
}

func update(msg Msg, model Model_R) rt.SkyTuple2 {
	return rt.Coerce[rt.SkyTuple2](func() any { __subject := any(msg).(Msg); if __subject.Tag == 0 {; 	return func() any { __subject := rt.Field(model, "SourceMode"); if rt.EnumTagIs(__subject, 0) {; 	return func() any { result := collectFileEntries(rt.AsListT[WatchedFile_R](rt.Field(model, "Watched")), rt.AsMapT[int](rt.Field(model, "FileCounts")), rt.AsListT[Log_Entry_LogEntry_R](rt.Field(model, "Entries"))); _ = result; return func() any { _ = rt.AnyTaskRun(Log_Webhook_processNewEntries(rt.AsListT[Log_Webhook_WebhookRule_R](rt.Field(model, "WebhookRules")), rt.AsListT[Log_Entry_LogEntry_R](rt.Field(result, "NewEntries")))); return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"Entries": rt.Field(result, "Entries"), "FileCounts": rt.Field(result, "Counts")}), V1: rt.Cmd_none()} }() }(); }; if rt.EnumTagIs(__subject, 1) {; 	return func() any { newEntries := collectStreamEntries(rt.AsListT[SourceScanner_R](rt.Field(model, "Scanners"))); _ = newEntries; return func() any { _ = rt.AnyTaskRun(Log_Webhook_processNewEntries(rt.AsListT[Log_Webhook_WebhookRule_R](rt.Field(model, "WebhookRules")), rt.AsListT[Log_Entry_LogEntry_R](newEntries))); return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"Entries": rt.List_append(rt.Field(model, "Entries"), newEntries)}), V1: rt.Cmd_none()} }() }(); }; _ = rt.Unreachable("case/__subject"); return nil }(); }; if __subject.Tag == 1 {; 	value := rt.AdtField(any(__subject), 0); 	_ = value; 	return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"LevelFilter": value}), V1: rt.Cmd_none()}; }; if __subject.Tag == 2 {; 	value := rt.AdtField(any(__subject), 0); 	_ = value; 	return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"ScopeFilter": value}), V1: rt.Cmd_none()}; }; if __subject.Tag == 3 {; 	value := rt.AdtField(any(__subject), 0); 	_ = value; 	return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"SearchFilter": value}), V1: rt.Cmd_none()}; }; if __subject.Tag == 4 {; 	value := rt.AdtField(any(__subject), 0); 	_ = value; 	return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"SourceFilter": value}), V1: rt.Cmd_none()}; }; if __subject.Tag == 5 {; 	return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"AutoScroll": rt.Basics_notT(rt.AsBool(rt.Field(model, "AutoScroll")))}), V1: rt.Cmd_none()}; }; if __subject.Tag == 6 {; 	return func() any { newTheme := func() any { if rt.AsBool(rt.Eq(rt.Field(model, "Theme"), "dark")) {; 	return "light"; } else {; 	return "dark"; }; return nil }(); _ = newTheme; return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"Theme": newTheme}), V1: rt.Cmd_none()} }(); }; if __subject.Tag == 7 {; 	return rt.SkyTuple2{V0: rt.RecordUpdate(model, map[string]any{"Entries": []any{}}), V1: rt.Cmd_none()}; }; if __subject.Tag == 8 {; 	_ = rt.AdtField(any(__subject), 0); 	return rt.SkyTuple2{V0: model, V1: rt.Cmd_none()}; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func subscriptions(_ Model_R) rt.SkySub {
	return rt.Coerce[rt.SkySub](rt.Time_every(200, Msg_Tick))
}

func filterEntry(levelStr string, scopeStr string, searchStr string, sourceStr string, entry Log_Entry_LogEntry_R) bool {
	return rt.CoerceBool(func() any { levelOk := func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(levelStr))) {; 	return true; } else {; 	return rt.Gte(Log_Entry_levelToInt(any(rt.Field(entry, "Level")).(Log_Entry_Level)), Log_Entry_levelToInt(any(Log_Entry_levelFromString(rt.CoerceString(levelStr))).(Log_Entry_Level))); }; return nil }(); _ = levelOk; return func() any { scopeOk := func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(scopeStr))) {; 	return true; } else {; 	return rt.String_containsT(rt.AsString(rt.String_toLowerT(rt.AsString(scopeStr))), rt.AsString(rt.String_toLowerT(rt.AsString(rt.Field(entry, "Scope"))))); }; return nil }(); _ = scopeOk; return func() any { searchOk := func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(searchStr))) {; 	return true; } else {; 	return func() any { __subject_tFfi := rt.Go_Regexp_compileT(rt.CoerceString(rt.Concat("(?i)", searchStr))); if __subject_tFfi.Tag == 0 {; 	re := any(__subject_tFfi.OkValue); 	_ = re; 	return rt.Result_withDefaultAnyT(any(false), any(rt.Go_Regexp_regexpMatchStringT(rt.Coerce[rt.FfiT_Go_Regexp_regexpMatchString_P0](re), rt.CoerceString(rt.Field(entry, "Message"))))); }; if __subject_tFfi.Tag == 1 {; 	_ = any(__subject_tFfi.ErrValue); 	return rt.String_containsT(rt.AsString(rt.String_toLowerT(rt.AsString(searchStr))), rt.AsString(rt.String_toLowerT(rt.AsString(rt.Field(entry, "Message"))))); }; _ = rt.Unreachable("case/__subject_tFfi"); return nil }(); }; return nil }(); _ = searchOk; return func() any { sourceOk := func() any { if rt.AsBool(rt.String_isEmptyT(rt.AsString(sourceStr))) {; 	return true; } else {; 	return rt.Eq(rt.Field(entry, "Source"), sourceStr); }; return nil }(); _ = sourceOk; return rt.And(levelOk, rt.And(scopeOk, rt.And(searchOk, sourceOk))) }() }() }() }())
}

func levelClass(level Log_Entry_Level) string {
	return rt.CoerceString(func() any { __subject := level; if rt.EnumTagIs(__subject, 0) {; 	return "log-entry l-debug"; }; if rt.EnumTagIs(__subject, 1) {; 	return "log-entry l-info"; }; if rt.EnumTagIs(__subject, 2) {; 	return "log-entry l-warn"; }; if rt.EnumTagIs(__subject, 3) {; 	return "log-entry l-error"; }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func padRight(targetLen int, pad string, str string) string {
	return rt.CoerceString(func() any { if rt.AsBool(rt.Gte(rt.String_lengthT(rt.AsString(str)), targetLen)) {; 	return str; } else {; 	return padRight(rt.CoerceInt(targetLen), rt.CoerceString(pad), rt.CoerceString(rt.Concat(str, pad))); }; return nil }())
}

func formatTimestamp(ts int) string {
	return rt.CoerceString(func() any { secs := func() any { if rt.AsBool(rt.Gt(ts, 9999999999)) {; 	return rt.IntDiv(ts, 1000); } else {; 	return ts; }; return nil }(); _ = secs; return func() any { totalSeconds := rt.Sub(secs, rt.Mul(rt.IntDiv(secs, 86400), 86400)); _ = totalSeconds; return func() any { h := rt.IntDiv(totalSeconds, 3600); _ = h; return func() any { remainder := rt.Sub(totalSeconds, rt.Mul(h, 3600)); _ = remainder; return func() any { m := rt.IntDiv(remainder, 60); _ = m; return func() any { s := rt.Sub(remainder, rt.Mul(m, 60)); _ = s; return rt.Concat(padZero(rt.CoerceInt(h)), rt.Concat(":", rt.Concat(padZero(rt.CoerceInt(m)), rt.Concat(":", padZero(rt.CoerceInt(s)))))) }() }() }() }() }() }())
}

func padZero(n int) string {
	return rt.CoerceString(func() any { if rt.AsBool(rt.Lt(n, 10)) {; 	return rt.Concat("0", rt.String_fromIntT(rt.AsInt(n))); } else {; 	return rt.String_fromIntT(rt.AsInt(n)); }; return nil }())
}

func viewEntry(entry Log_Entry_LogEntry_R) rt.VNode {
	return rt.Coerce[rt.VNode](rt.Html_div([]any{rt.Attr_classT(rt.AsString(levelClass(any(rt.Field(entry, "Level")).(Log_Entry_Level))))}, []any{rt.Html_span([]any{rt.Attr_classT("ts")}, []any{rt.Html_textT(rt.AsString(formatTimestamp(rt.CoerceInt(rt.Field(entry, "Timestamp")))))}), rt.Html_span([]any{rt.Attr_classT("level")}, []any{rt.Html_textT(rt.AsString(rt.String_toUpperT(rt.AsString(padRight(rt.CoerceInt(5), rt.CoerceString(" "), rt.CoerceString(Log_Entry_levelToString(any(rt.Field(entry, "Level")).(Log_Entry_Level))))))))}), rt.Html_span([]any{rt.Attr_classT("source")}, []any{rt.Html_textT(rt.AsString(rt.Field(entry, "Source")))}), rt.Html_span([]any{rt.Attr_classT("scope")}, []any{rt.Html_textT(rt.AsString(rt.Field(entry, "Scope")))}), rt.Html_span([]any{rt.Attr_classT("msg")}, []any{rt.Html_textT(rt.AsString(rt.Field(entry, "Message")))})}))
}

func viewEntries(entries []Log_Entry_LogEntry_R) rt.VNode {
	return rt.Coerce[rt.VNode](func() any { if rt.AsBool(rt.List_isEmptyT(rt.AsList(entries))) {; 	return rt.Html_div([]any{rt.Attr_classT("empty")}, []any{rt.Html_textT("Waiting for logs...")}); } else {; 	return rt.Html_div([]any{rt.Attr_id("log-entries")}, rt.List_mapAny(any(viewEntry), any(rt.List_takeAnyT(5000, rt.AsList(entries))))); }; return nil }())
}

func statusText(model Model_R) string {
	return rt.CoerceString(func() any { __subject := rt.Field(model, "SourceMode"); if rt.EnumTagIs(__subject, 0) {; 	return func() any { labels := rt.List_mapAny(any(func(w any) any { return rt.Field(w, "Label"); }), any(rt.Field(model, "Watched"))); _ = labels; return rt.Concat("Watching ", rt.Concat(rt.String_fromIntT(rt.AsInt(rt.List_lengthT(rt.AsList(rt.Field(model, "Watched"))))), rt.Concat(" source(s): ", rt.String_join(", ", labels)))) }(); }; if rt.EnumTagIs(__subject, 1) {; 	return func() any { labels := rt.List_mapAny(any(func(s any) any { return rt.Field(s, "Label"); }), any(rt.Field(model, "Scanners"))); _ = labels; return rt.Concat("Streaming: ", rt.String_join(", ", labels)) }(); }; _ = rt.Unreachable("case/__subject"); return nil }())
}

func styles() rt.VNode {
	return rt.Coerce[rt.VNode](rt.Html_styleNode([]any{}, rt.Css_stylesheet([]any{rt.Css_rule("*", []any{rt.Css_margin(rt.Css_zero(struct{}{})), rt.Css_padding(rt.Css_zero(struct{}{})), rt.Css_propertyT("box-sizing", "border-box")}), rt.Css_rule("html, body", []any{rt.Css_height(rt.Css_pct(100)), rt.Css_overflow("hidden")}), rt.Css_rule(".root.dark", []any{rt.Css_propertyT("--bg", "#0a0e14"), rt.Css_propertyT("--fg", "#c5cdd8"), rt.Css_propertyT("--muted", "#64748b"), rt.Css_propertyT("--dimmed", "#475569"), rt.Css_propertyT("--input-bg", "#0a0e14"), rt.Css_propertyT("--input-border", "rgba(255,255,255,0.12)"), rt.Css_propertyT("--border", "rgba(255,255,255,0.06)"), rt.Css_propertyT("--hover", "rgba(255,255,255,0.02)"), rt.Css_propertyT("--scroll-thumb", "rgba(255,255,255,0.08)"), rt.Css_propertyT("--scroll-hover", "rgba(255,255,255,0.15)"), rt.Css_propertyT("--btn-border", "rgba(255,255,255,0.08)"), rt.Css_propertyT("--btn-hover-fg", "#c5cdd8"), rt.Css_propertyT("--btn-hover-border", "rgba(255,255,255,0.15)")}), rt.Css_rule(".root.light", []any{rt.Css_propertyT("--bg", "#ffffff"), rt.Css_propertyT("--fg", "#1e293b"), rt.Css_propertyT("--muted", "#94a3b8"), rt.Css_propertyT("--dimmed", "#94a3b8"), rt.Css_propertyT("--input-bg", "#f8fafc"), rt.Css_propertyT("--input-border", "#d1d5db"), rt.Css_propertyT("--border", "#e5e7eb"), rt.Css_propertyT("--hover", "rgba(0,0,0,0.02)"), rt.Css_propertyT("--scroll-thumb", "rgba(0,0,0,0.12)"), rt.Css_propertyT("--scroll-hover", "rgba(0,0,0,0.2)"), rt.Css_propertyT("--btn-border", "#d1d5db"), rt.Css_propertyT("--btn-hover-fg", "#1e293b"), rt.Css_propertyT("--btn-hover-border", "#9ca3af")}), rt.Css_rule("body", []any{rt.Css_backgroundColor(rt.Css_hexT("#0a0e14")), rt.Css_color(rt.Css_hexT("#c5cdd8")), rt.Css_fontFamily("'SF Mono', 'JetBrains Mono', ui-monospace, Menlo, monospace"), rt.Css_fontSize(rt.Css_pxT(11)), rt.Css_propertyT("-webkit-font-smoothing", "antialiased")}), rt.Css_rule(".root", []any{rt.Css_position("fixed"), rt.Css_top(rt.Css_zero(struct{}{})), rt.Css_propertyT("left", "0"), rt.Css_propertyT("right", "0"), rt.Css_bottom(rt.Css_zero(struct{}{})), rt.Css_display("flex"), rt.Css_flexDirection("column"), rt.Css_overflow("hidden"), rt.Css_backgroundColor("var(--bg, #0a0e14)"), rt.Css_color("var(--fg, #c5cdd8)")}), rt.Css_rule(".toolbar", []any{rt.Css_display("flex"), rt.Css_alignItems("center"), rt.Css_gap(rt.Css_pxT(8)), rt.Css_padding2(rt.Css_pxT(6), rt.Css_pxT(12)), rt.Css_propertyT("border-bottom", "1px solid var(--border)"), rt.Css_propertyT("flex-shrink", "0")}), rt.Css_rule(".toolbar h1", []any{rt.Css_fontSize(rt.Css_pxT(11)), rt.Css_fontWeight("500"), rt.Css_color("var(--muted)"), rt.Css_propertyT("margin-right", "4px")}), rt.Css_rule(".toolbar h1 span", []any{rt.Css_color("var(--fg)"), rt.Css_fontWeight("600")}), rt.Css_rule(".toolbar select, .toolbar input", []any{rt.Css_backgroundColor("var(--input-bg)"), rt.Css_color("var(--fg)"), rt.Css_propertyT("border", "1px solid var(--input-border)"), rt.Css_borderRadius(rt.Css_pxT(4)), rt.Css_padding2(rt.Css_pxT(3), rt.Css_pxT(6)), rt.Css_fontFamily("inherit"), rt.Css_fontSize(rt.Css_pxT(11)), rt.Css_propertyT("outline", "none"), rt.Css_transition("border-color 0.15s")}), rt.Css_rule(".toolbar select:focus, .toolbar input:focus", []any{rt.Css_borderColor(rt.Css_hexT("#60a5fa"))}), rt.Css_rule(".toolbar input", []any{rt.Css_width(rt.Css_pxT(120))}), rt.Css_rule(".spacer", []any{rt.Css_flex("1")}), rt.Css_rule(".stats", []any{rt.Css_color("var(--dimmed)"), rt.Css_fontSize(rt.Css_pxT(11))}), rt.Css_rule(".btn", []any{rt.Css_backgroundColor("transparent"), rt.Css_color("var(--muted)"), rt.Css_propertyT("border", "1px solid var(--btn-border)"), rt.Css_borderRadius(rt.Css_pxT(4)), rt.Css_padding2(rt.Css_pxT(3), rt.Css_pxT(8)), rt.Css_fontFamily("inherit"), rt.Css_fontSize(rt.Css_pxT(11)), rt.Css_cursor("pointer"), rt.Css_transition("all 0.15s"), rt.Css_propertyT("user-select", "none")}), rt.Css_rule(".btn:hover", []any{rt.Css_color("var(--btn-hover-fg)"), rt.Css_propertyT("border-color", "var(--btn-hover-border)")}), rt.Css_rule(".btn.active", []any{rt.Css_color(rt.Css_hexT("#60a5fa")), rt.Css_borderColor("rgba(96,165,250,0.3)")}), rt.Css_rule(".btn.danger:hover", []any{rt.Css_color(rt.Css_hexT("#f87171")), rt.Css_borderColor("rgba(248,113,113,0.3)")}), rt.Css_rule("#log-entries", []any{rt.Css_flex("1"), rt.Css_propertyT("min-height", "0"), rt.Css_propertyT("overflow-y", "scroll")}), rt.Css_rule("#log-entries::-webkit-scrollbar", []any{rt.Css_width(rt.Css_pxT(6))}), rt.Css_rule("#log-entries::-webkit-scrollbar-track", []any{rt.Css_backgroundColor("transparent")}), rt.Css_rule("#log-entries::-webkit-scrollbar-thumb", []any{rt.Css_propertyT("background-color", "var(--scroll-thumb)"), rt.Css_borderRadius(rt.Css_pxT(3))}), rt.Css_rule("#log-entries::-webkit-scrollbar-thumb:hover", []any{rt.Css_propertyT("background-color", "var(--scroll-hover)")}), rt.Css_rule(".log-entry", []any{rt.Css_display("flex"), rt.Css_alignItems("baseline"), rt.Css_padding2(rt.Css_pxT(0), rt.Css_pxT(12)), rt.Css_propertyT("line-height", "18px"), rt.Css_gap(rt.Css_pxT(8)), rt.Css_borderLeft("2px solid transparent")}), rt.Css_rule(".log-entry:hover", []any{rt.Css_backgroundColor("var(--hover)")}), rt.Css_rule(".l-debug .level", []any{rt.Css_color(rt.Css_hexT("#64748b"))}), rt.Css_rule(".l-info .level", []any{rt.Css_color(rt.Css_hexT("#60a5fa"))}), rt.Css_rule(".l-warn .level", []any{rt.Css_color(rt.Css_hexT("#fbbf24"))}), rt.Css_rule(".l-error .level", []any{rt.Css_color(rt.Css_hexT("#f87171"))}), rt.Css_rule(".log-entry.l-error", []any{rt.Css_borderLeft("2px solid rgba(248,113,113,0.5)"), rt.Css_backgroundColor("rgba(248,113,113,0.03)")}), rt.Css_rule(".log-entry.l-warn", []any{rt.Css_borderLeft("2px solid rgba(251,191,36,0.4)"), rt.Css_backgroundColor("rgba(251,191,36,0.02)")}), rt.Css_rule(".ts", []any{rt.Css_color("var(--dimmed)"), rt.Css_propertyT("white-space", "nowrap"), rt.Css_minWidth(rt.Css_pxT(56)), rt.Css_propertyT("font-variant-numeric", "tabular-nums")}), rt.Css_rule(".level", []any{rt.Css_fontWeight("600"), rt.Css_propertyT("white-space", "nowrap"), rt.Css_minWidth(rt.Css_pxT(40)), rt.Css_textTransform("uppercase"), rt.Css_letterSpacing("0.3px")}), rt.Css_rule(".source", []any{rt.Css_propertyT("white-space", "nowrap"), rt.Css_maxWidth(rt.Css_pxT(100)), rt.Css_overflow("hidden"), rt.Css_propertyT("text-overflow", "ellipsis"), rt.Css_color(rt.Css_hexT("#a78bfa")), rt.Css_backgroundColor("rgba(167,139,250,0.08)"), rt.Css_padding2(rt.Css_pxT(0), rt.Css_pxT(4)), rt.Css_borderRadius(rt.Css_pxT(3))}), rt.Css_rule(".scope", []any{rt.Css_color(rt.Css_hexT("#34d399")), rt.Css_propertyT("white-space", "nowrap"), rt.Css_maxWidth(rt.Css_pxT(120)), rt.Css_overflow("hidden"), rt.Css_propertyT("text-overflow", "ellipsis")}), rt.Css_rule(".msg", []any{rt.Css_color("var(--fg)"), rt.Css_propertyT("word-break", "break-word"), rt.Css_flex("1")}), rt.Css_rule(".empty", []any{rt.Css_display("flex"), rt.Css_alignItems("center"), rt.Css_justifyContent("center"), rt.Css_height(rt.Css_pct(100)), rt.Css_color("var(--dimmed)")}), rt.Css_rule(".status", []any{rt.Css_display("flex"), rt.Css_alignItems("center"), rt.Css_gap(rt.Css_pxT(6)), rt.Css_padding2(rt.Css_pxT(4), rt.Css_pxT(12)), rt.Css_propertyT("border-top", "1px solid var(--border)"), rt.Css_fontSize(rt.Css_pxT(10)), rt.Css_color("var(--dimmed)"), rt.Css_propertyT("flex-shrink", "0")}), rt.Css_rule(".dot", []any{rt.Css_width(rt.Css_pxT(5)), rt.Css_height(rt.Css_pxT(5)), rt.Css_borderRadius(rt.Css_pct(50)), rt.Css_backgroundColor(rt.Css_hexT("#34d399")), rt.Css_propertyT("box-shadow", "0 0 4px rgba(52,211,153,0.4)")})})))
}

func uniqueSources(entries []Log_Entry_LogEntry_R) []string {
	return rt.Coerce[[]string](rt.List_foldlAnyT(any(func(s any) any { return func(acc any) any { return func() any { if rt.AsBool(rt.List_member(s, acc)) {; 	return acc; } else {; 	return rt.List_append(acc, []any{s}); }; return nil }(); }; }), any([]any{}), rt.AsList(rt.List_filterAny(any(func(s any) any { return rt.Basics_notT(rt.AsBool(rt.String_isEmptyT(rt.AsString(s)))); }), any(rt.List_mapAny(any(func(e any) any { return rt.Field(e, "Source"); }), any(entries)))))))
}

func view(model Model_R) rt.VNode {
	return rt.Coerce[rt.VNode](func() any { filtered := rt.List_filterAny(any(func(__pp0 any) any { return filterEntry(rt.CoerceString(rt.Field(model, "LevelFilter")), rt.CoerceString(rt.Field(model, "ScopeFilter")), rt.CoerceString(rt.Field(model, "SearchFilter")), rt.CoerceString(rt.Field(model, "SourceFilter")), any(__pp0).(Log_Entry_LogEntry_R)); }), any(rt.Field(model, "Entries"))); _ = filtered; return func() any { sources := uniqueSources(rt.AsListT[Log_Entry_LogEntry_R](rt.Field(model, "Entries"))); _ = sources; return func() any { totalCount := rt.List_lengthT(rt.AsList(rt.Field(model, "Entries"))); _ = totalCount; return func() any { filteredCount := rt.List_lengthT(rt.AsList(filtered)); _ = filteredCount; return func() any { scrollLabel := func() any { if rt.AsBool(rt.Field(model, "AutoScroll")) {; 	return "Auto-scroll"; } else {; 	return "Paused"; }; return nil }(); _ = scrollLabel; return func() any { scrollClass := func() any { if rt.AsBool(rt.Field(model, "AutoScroll")) {; 	return "btn active"; } else {; 	return "btn"; }; return nil }(); _ = scrollClass; return func() any { themeIcon := func() any { if rt.AsBool(rt.Eq(rt.Field(model, "Theme"), "dark")) {; 	return "☼"; } else {; 	return "☾"; }; return nil }(); _ = themeIcon; return rt.Html_div([]any{rt.Attr_classT(rt.AsString(rt.Concat("root ", rt.Field(model, "Theme"))))}, []any{styles(), rt.Html_div([]any{rt.Attr_classT("toolbar")}, []any{rt.Html_h1([]any{}, []any{rt.Html_textT("sky"), rt.Html_span([]any{}, []any{rt.Html_textT("-log")})}), rt.Html_button([]any{rt.Attr_classT("btn"), rt.Event_onClick(Msg_ToggleTheme)}, []any{rt.Html_textT(rt.AsString(themeIcon))}), rt.Html_select([]any{rt.Attr_id("level-filter"), rt.Event_onChange(Msg_SetLevelFilter)}, []any{rt.Html_option([]any{rt.Attr_value("")}, []any{rt.Html_textT("Level")}), rt.Html_option([]any{rt.Attr_value("debug")}, []any{rt.Html_textT("DEBUG")}), rt.Html_option([]any{rt.Attr_value("info")}, []any{rt.Html_textT("INFO")}), rt.Html_option([]any{rt.Attr_value("warn")}, []any{rt.Html_textT("WARN")}), rt.Html_option([]any{rt.Attr_value("error")}, []any{rt.Html_textT("ERROR")})}), rt.Html_input([]any{rt.Attr_id("scope-filter"), rt.Attr_type("text"), rt.Attr_placeholder("scope"), rt.Event_onInput(Msg_SetScopeFilter)}), rt.Html_input([]any{rt.Attr_id("search-filter"), rt.Attr_type("text"), rt.Attr_placeholder("regex search"), rt.Event_onInput(Msg_SetSearchFilter)}), rt.Html_select([]any{rt.Attr_id("source-filter"), rt.Event_onChange(Msg_SetSourceFilter)}, rt.List_append([]any{rt.Html_option([]any{rt.Attr_value("")}, []any{rt.Html_textT("Source")})}, rt.List_mapAny(any(func(s any) any { return rt.Html_option([]any{rt.Attr_value(s)}, []any{rt.Html_textT(rt.AsString(s))}); }), any(sources)))), rt.Html_div([]any{rt.Attr_classT("spacer")}, []any{}), rt.Html_button([]any{rt.Attr_classT(rt.AsString(scrollClass)), rt.Event_onClick(Msg_ToggleAutoScroll)}, []any{rt.Html_textT(rt.AsString(scrollLabel))}), rt.Html_button([]any{rt.Attr_classT("btn danger"), rt.Event_onClick(Msg_ClearLogs)}, []any{rt.Html_textT("Clear")})}), viewEntries(rt.AsListT[Log_Entry_LogEntry_R](filtered)), rt.Html_div([]any{rt.Attr_classT("status")}, []any{rt.Html_span([]any{rt.Attr_classT("dot")}, []any{}), rt.Html_textT(rt.AsString(statusText(any(model).(Model_R)))), rt.Html_div([]any{rt.Attr_classT("spacer")}, []any{}), rt.Html_span([]any{rt.Attr_classT("stats")}, []any{rt.Html_textT(rt.AsString(rt.Concat(rt.String_fromIntT(rt.AsInt(filteredCount)), rt.Concat("/", rt.String_fromIntT(rt.AsInt(totalCount))))))})})}) }() }() }() }() }() }() }())
}

func main() {
	_ = rt.AnyTaskRun(rt.Live_app(struct{ Init any; NotFound any; Routes any; Subscriptions any; Update any; View any }{Init: init_[any], NotFound: Page_LogPage, Routes: []any{rt.Live_route("/", Page_LogPage)}, Subscriptions: subscriptions, Update: update, View: view}))
}

