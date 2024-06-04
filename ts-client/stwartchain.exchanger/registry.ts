import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgExchange } from "./types/stwartchain/exchanger/tx";
import { MsgUpdateParams } from "./types/stwartchain/exchanger/tx";
import { MsgUpdateParamsResponse } from "./types/stwartchain/exchanger/tx";
import { Params } from "./types/stwartchain/exchanger/params";
import { QueryParamsResponse } from "./types/stwartchain/exchanger/query";
import { MsgExchangeResponse } from "./types/stwartchain/exchanger/tx";
import { GenesisState } from "./types/stwartchain/exchanger/genesis";
import { QueryParamsRequest } from "./types/stwartchain/exchanger/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.exchanger.MsgExchange", MsgExchange],
    ["/stwartchain.exchanger.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.exchanger.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.exchanger.Params", Params],
    ["/stwartchain.exchanger.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.exchanger.MsgExchangeResponse", MsgExchangeResponse],
    ["/stwartchain.exchanger.GenesisState", GenesisState],
    ["/stwartchain.exchanger.QueryParamsRequest", QueryParamsRequest],
    
];

export { msgTypes }