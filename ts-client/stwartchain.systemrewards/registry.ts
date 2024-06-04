import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryAllStatsResponse } from "./types/stwartchain/systemrewards/query";
import { MsgUpdateParamsResponse } from "./types/stwartchain/systemrewards/tx";
import { Params } from "./types/stwartchain/systemrewards/params";
import { GenesisState } from "./types/stwartchain/systemrewards/genesis";
import { QueryParamsResponse } from "./types/stwartchain/systemrewards/query";
import { QueryAllStatsRequest } from "./types/stwartchain/systemrewards/query";
import { QueryStatsByDateRequest } from "./types/stwartchain/systemrewards/query";
import { DailyStats } from "./types/stwartchain/systemrewards/daily_stats";
import { QueryGetStatsRequest } from "./types/stwartchain/systemrewards/query";
import { QueryStatsByDateResponse } from "./types/stwartchain/systemrewards/query";
import { QueryParamsRequest } from "./types/stwartchain/systemrewards/query";
import { QueryGetStatsResponse } from "./types/stwartchain/systemrewards/query";
import { MsgUpdateParams } from "./types/stwartchain/systemrewards/tx";
import { Stats } from "./types/stwartchain/systemrewards/stats";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.systemrewards.QueryAllStatsResponse", QueryAllStatsResponse],
    ["/stwartchain.systemrewards.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.systemrewards.Params", Params],
    ["/stwartchain.systemrewards.GenesisState", GenesisState],
    ["/stwartchain.systemrewards.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.systemrewards.QueryAllStatsRequest", QueryAllStatsRequest],
    ["/stwartchain.systemrewards.QueryStatsByDateRequest", QueryStatsByDateRequest],
    ["/stwartchain.systemrewards.DailyStats", DailyStats],
    ["/stwartchain.systemrewards.QueryGetStatsRequest", QueryGetStatsRequest],
    ["/stwartchain.systemrewards.QueryStatsByDateResponse", QueryStatsByDateResponse],
    ["/stwartchain.systemrewards.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.systemrewards.QueryGetStatsResponse", QueryGetStatsResponse],
    ["/stwartchain.systemrewards.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.systemrewards.Stats", Stats],
    
];

export { msgTypes }