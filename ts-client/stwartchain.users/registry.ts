import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryAllStatsResponse } from "./types/stwartchain/users/query";
import { QueryGetUniqueUsersResponse } from "./types/stwartchain/users/query";
import { MsgUpdateParams } from "./types/stwartchain/users/tx";
import { Stats } from "./types/stwartchain/users/stats";
import { QueryStatsByDateResponse } from "./types/stwartchain/users/query";
import { QueryGetStatsResponse } from "./types/stwartchain/users/query";
import { Params } from "./types/stwartchain/users/params";
import { UniqueUserAddresses } from "./types/stwartchain/users/unique_user_addresses";
import { QueryAllStatsRequest } from "./types/stwartchain/users/query";
import { QueryStatsByDateRequest } from "./types/stwartchain/users/query";
import { QueryAllUniqueUsersRequest } from "./types/stwartchain/users/query";
import { DailyStats } from "./types/stwartchain/users/daily_stats";
import { UniqueUsers } from "./types/stwartchain/users/unique_users";
import { QueryTotalRequest } from "./types/stwartchain/users/query";
import { GenesisState } from "./types/stwartchain/users/genesis";
import { QueryParamsResponse } from "./types/stwartchain/users/query";
import { QueryGetStatsRequest } from "./types/stwartchain/users/query";
import { TotalUsers } from "./types/stwartchain/users/total_users";
import { QueryParamsRequest } from "./types/stwartchain/users/query";
import { QueryGetUniqueUsersRequest } from "./types/stwartchain/users/query";
import { QueryAllUniqueUsersResponse } from "./types/stwartchain/users/query";
import { QueryTotalResponse } from "./types/stwartchain/users/query";
import { MsgUpdateParamsResponse } from "./types/stwartchain/users/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.users.QueryAllStatsResponse", QueryAllStatsResponse],
    ["/stwartchain.users.QueryGetUniqueUsersResponse", QueryGetUniqueUsersResponse],
    ["/stwartchain.users.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.users.Stats", Stats],
    ["/stwartchain.users.QueryStatsByDateResponse", QueryStatsByDateResponse],
    ["/stwartchain.users.QueryGetStatsResponse", QueryGetStatsResponse],
    ["/stwartchain.users.Params", Params],
    ["/stwartchain.users.UniqueUserAddresses", UniqueUserAddresses],
    ["/stwartchain.users.QueryAllStatsRequest", QueryAllStatsRequest],
    ["/stwartchain.users.QueryStatsByDateRequest", QueryStatsByDateRequest],
    ["/stwartchain.users.QueryAllUniqueUsersRequest", QueryAllUniqueUsersRequest],
    ["/stwartchain.users.DailyStats", DailyStats],
    ["/stwartchain.users.UniqueUsers", UniqueUsers],
    ["/stwartchain.users.QueryTotalRequest", QueryTotalRequest],
    ["/stwartchain.users.GenesisState", GenesisState],
    ["/stwartchain.users.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.users.QueryGetStatsRequest", QueryGetStatsRequest],
    ["/stwartchain.users.TotalUsers", TotalUsers],
    ["/stwartchain.users.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.users.QueryGetUniqueUsersRequest", QueryGetUniqueUsersRequest],
    ["/stwartchain.users.QueryAllUniqueUsersResponse", QueryAllUniqueUsersResponse],
    ["/stwartchain.users.QueryTotalResponse", QueryTotalResponse],
    ["/stwartchain.users.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    
];

export { msgTypes }