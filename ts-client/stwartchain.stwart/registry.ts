import { GeneratedType } from "@cosmjs/proto-signing";
import { Params } from "./types/stwartchain/stwart/params";
import { QueryParamsRequest } from "./types/stwartchain/stwart/query";
import { QueryParamsResponse } from "./types/stwartchain/stwart/query";
import { MsgUpdateParamsResponse } from "./types/stwartchain/stwart/tx";
import { MsgUpdateParams } from "./types/stwartchain/stwart/tx";
import { GenesisState } from "./types/stwartchain/stwart/genesis";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.stwart.Params", Params],
    ["/stwartchain.stwart.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.stwart.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.stwart.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.stwart.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.stwart.GenesisState", GenesisState],
    
];

export { msgTypes }