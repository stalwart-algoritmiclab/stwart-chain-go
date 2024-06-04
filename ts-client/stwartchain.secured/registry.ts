import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCreateAddressesResponse } from "./types/stwartchain/secured/tx";
import { MsgUpdateAddresses } from "./types/stwartchain/secured/tx";
import { MsgDeleteByAddresses } from "./types/stwartchain/secured/tx";
import { Params } from "./types/stwartchain/secured/params";
import { QueryGetAddressesResponse } from "./types/stwartchain/secured/query";
import { MsgUpdateParams } from "./types/stwartchain/secured/tx";
import { QueryAllAddressesResponse } from "./types/stwartchain/secured/query";
import { QueryGetAddressRequest } from "./types/stwartchain/secured/query";
import { QueryAllAddressesRequest } from "./types/stwartchain/secured/query";
import { QueryParamsRequest } from "./types/stwartchain/secured/query";
import { QueryParamsResponse } from "./types/stwartchain/secured/query";
import { GenesisState } from "./types/stwartchain/secured/genesis";
import { MsgUpdateParamsResponse } from "./types/stwartchain/secured/tx";
import { MsgCreateAddresses } from "./types/stwartchain/secured/tx";
import { MsgUpdateAddressesResponse } from "./types/stwartchain/secured/tx";
import { MsgDeleteAddresses } from "./types/stwartchain/secured/tx";
import { MsgDeleteAddressesResponse } from "./types/stwartchain/secured/tx";
import { Addresses } from "./types/stwartchain/secured/addresses";
import { QueryGetAddressesRequest } from "./types/stwartchain/secured/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/stwartchain.secured.MsgCreateAddressesResponse", MsgCreateAddressesResponse],
    ["/stwartchain.secured.MsgUpdateAddresses", MsgUpdateAddresses],
    ["/stwartchain.secured.MsgDeleteByAddresses", MsgDeleteByAddresses],
    ["/stwartchain.secured.Params", Params],
    ["/stwartchain.secured.QueryGetAddressesResponse", QueryGetAddressesResponse],
    ["/stwartchain.secured.MsgUpdateParams", MsgUpdateParams],
    ["/stwartchain.secured.QueryAllAddressesResponse", QueryAllAddressesResponse],
    ["/stwartchain.secured.QueryGetAddressRequest", QueryGetAddressRequest],
    ["/stwartchain.secured.QueryAllAddressesRequest", QueryAllAddressesRequest],
    ["/stwartchain.secured.QueryParamsRequest", QueryParamsRequest],
    ["/stwartchain.secured.QueryParamsResponse", QueryParamsResponse],
    ["/stwartchain.secured.GenesisState", GenesisState],
    ["/stwartchain.secured.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/stwartchain.secured.MsgCreateAddresses", MsgCreateAddresses],
    ["/stwartchain.secured.MsgUpdateAddressesResponse", MsgUpdateAddressesResponse],
    ["/stwartchain.secured.MsgDeleteAddresses", MsgDeleteAddresses],
    ["/stwartchain.secured.MsgDeleteAddressesResponse", MsgDeleteAddressesResponse],
    ["/stwartchain.secured.Addresses", Addresses],
    ["/stwartchain.secured.QueryGetAddressesRequest", QueryGetAddressesRequest],
    
];

export { msgTypes }