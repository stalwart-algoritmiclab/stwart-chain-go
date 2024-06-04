import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryGetFeeStatsResponse } from "./types/stwartchain/stats/query";
import { MsgUpdateParamsResponse } from "./types/stwartchain/stats/tx";
import { QueryDateRequest } from "./types/stwartchain/stats/query";
import { QueryUserStatsResponse } from "./types/stwartchain/stats/query";
import { UserStats } from "./types/stwartchain/stats/user_stats";
import { AssetStats } from "./types/stwartchain/stats/stats";
import { MsgUpdateParams } from "./types/stwartchain/stats/tx";
import { QueryDateWithPaginationRequest } from "./types/stwartchain/stats/query";
import { QueryAllFeeStatsRequest } from "./types/stwartchain/stats/query";
import { QueryGetFeeStatsByIndexesRequest } from "./types/stwartchain/stats/query";
import { Params } from "./types/stwartchain/stats/params";
import { QueryGetFeeStatsRequest } from "./types/stwartchain/stats/query";
import { FeeStats } from "./types/stwartchain/stats/fee_stats";
import { FeeDailyStats } from "./types/stwartchain/stats/fee_daily_stats";
import { QueryAssetStatsResponse } from "./types/stwartchain/stats/query";
import { QueryAllFeeStatsResponse } from "./types/stwartchain/stats/query";
import { QueryGetFeeStatsByDateRequest } from "./types/stwartchain/stats/query";
import { GenesisState } from "./types/stwartchain/stats/genesis";
import { QueryParamsRequest } from "./types/stwartchain/stats/query";
import { QueryParamsResponse } from "./types/stwartchain/stats/query";
import { AssetDailyStats } from "./types/stwartchain/stats/asset_stats";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.stats.QueryGetFeeStatsResponse", QueryGetFeeStatsResponse],
    ["/stwartchain.stats.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.stats.QueryDateRequest", QueryDateRequest],
    ["/stwartchain.stats.QueryUserStatsResponse", QueryUserStatsResponse],
    ["/stwartchain.stats.UserStats", UserStats],
    ["/stwartchain.stats.AssetStats", AssetStats],
    ["/stwartchain.stats.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.stats.QueryDateWithPaginationRequest", QueryDateWithPaginationRequest],
    ["/stwartchain.stats.QueryAllFeeStatsRequest", QueryAllFeeStatsRequest],
    ["/stwartchain.stats.QueryGetFeeStatsByIndexesRequest", QueryGetFeeStatsByIndexesRequest],
    ["/stwartchain.stats.Params", Params],
    ["/stwartchain.stats.QueryGetFeeStatsRequest", QueryGetFeeStatsRequest],
    ["/stwartchain.stats.FeeStats", FeeStats],
    ["/stwartchain.stats.FeeDailyStats", FeeDailyStats],
    ["/stwartchain.stats.QueryAssetStatsResponse", QueryAssetStatsResponse],
    ["/stwartchain.stats.QueryAllFeeStatsResponse", QueryAllFeeStatsResponse],
    ["/stwartchain.stats.QueryGetFeeStatsByDateRequest", QueryGetFeeStatsByDateRequest],
    ["/stwartchain.stats.GenesisState", GenesisState],
    ["/stwartchain.stats.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.stats.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.stats.AssetDailyStats", AssetDailyStats],
    
];

export { msgTypes }