/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { FeeStats } from "./fee_stats";
import { Params } from "./params";
import { AssetStats } from "./stats";
import { UserStats } from "./user_stats";

export const protobufPackage = "stwartchain.stats";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryDateRequest {
  startDate: string;
  endDate: string;
}

export interface QueryDateWithPaginationRequest {
  startDate: string;
  endDate: string;
  pagination: PageRequest | undefined;
}

export interface QueryAssetStatsResponse {
  stats: AssetStats | undefined;
}

export interface QueryGetFeeStatsRequest {
  date: string;
}

export interface QueryGetFeeStatsResponse {
  feeStats: FeeStats | undefined;
}

export interface QueryAllFeeStatsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllFeeStatsResponse {
  feeStats: FeeStats[];
  pagination: PageResponse | undefined;
}

export interface QueryGetFeeStatsByIndexesRequest {
  startIndex: string;
  endIndex: string;
  pagination: PageRequest | undefined;
}

export interface QueryGetFeeStatsByDateRequest {
  startDate: string;
  endDate: string;
  pagination: PageRequest | undefined;
}

export interface QueryUserStatsResponse {
  stats: UserStats | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(base?: I): QueryParamsRequest {
    return QueryParamsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(base?: I): QueryParamsResponse {
    return QueryParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryDateRequest(): QueryDateRequest {
  return { startDate: "", endDate: "" };
}

export const QueryDateRequest = {
  encode(message: QueryDateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.startDate !== "") {
      writer.uint32(10).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(18).string(message.endDate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDateRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.startDate = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.endDate = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryDateRequest {
    return {
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
    };
  },

  toJSON(message: QueryDateRequest): unknown {
    const obj: any = {};
    if (message.startDate !== "") {
      obj.startDate = message.startDate;
    }
    if (message.endDate !== "") {
      obj.endDate = message.endDate;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryDateRequest>, I>>(base?: I): QueryDateRequest {
    return QueryDateRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryDateRequest>, I>>(object: I): QueryDateRequest {
    const message = createBaseQueryDateRequest();
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    return message;
  },
};

function createBaseQueryDateWithPaginationRequest(): QueryDateWithPaginationRequest {
  return { startDate: "", endDate: "", pagination: undefined };
}

export const QueryDateWithPaginationRequest = {
  encode(message: QueryDateWithPaginationRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.startDate !== "") {
      writer.uint32(10).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(18).string(message.endDate);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDateWithPaginationRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDateWithPaginationRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.startDate = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.endDate = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryDateWithPaginationRequest {
    return {
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryDateWithPaginationRequest): unknown {
    const obj: any = {};
    if (message.startDate !== "") {
      obj.startDate = message.startDate;
    }
    if (message.endDate !== "") {
      obj.endDate = message.endDate;
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryDateWithPaginationRequest>, I>>(base?: I): QueryDateWithPaginationRequest {
    return QueryDateWithPaginationRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryDateWithPaginationRequest>, I>>(
    object: I,
  ): QueryDateWithPaginationRequest {
    const message = createBaseQueryDateWithPaginationRequest();
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAssetStatsResponse(): QueryAssetStatsResponse {
  return { stats: undefined };
}

export const QueryAssetStatsResponse = {
  encode(message: QueryAssetStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stats !== undefined) {
      AssetStats.encode(message.stats, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAssetStatsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAssetStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.stats = AssetStats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAssetStatsResponse {
    return { stats: isSet(object.stats) ? AssetStats.fromJSON(object.stats) : undefined };
  },

  toJSON(message: QueryAssetStatsResponse): unknown {
    const obj: any = {};
    if (message.stats !== undefined) {
      obj.stats = AssetStats.toJSON(message.stats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAssetStatsResponse>, I>>(base?: I): QueryAssetStatsResponse {
    return QueryAssetStatsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAssetStatsResponse>, I>>(object: I): QueryAssetStatsResponse {
    const message = createBaseQueryAssetStatsResponse();
    message.stats = (object.stats !== undefined && object.stats !== null)
      ? AssetStats.fromPartial(object.stats)
      : undefined;
    return message;
  },
};

function createBaseQueryGetFeeStatsRequest(): QueryGetFeeStatsRequest {
  return { date: "" };
}

export const QueryGetFeeStatsRequest = {
  encode(message: QueryGetFeeStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.date !== "") {
      writer.uint32(10).string(message.date);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetFeeStatsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetFeeStatsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.date = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeStatsRequest {
    return { date: isSet(object.date) ? String(object.date) : "" };
  },

  toJSON(message: QueryGetFeeStatsRequest): unknown {
    const obj: any = {};
    if (message.date !== "") {
      obj.date = message.date;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetFeeStatsRequest>, I>>(base?: I): QueryGetFeeStatsRequest {
    return QueryGetFeeStatsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetFeeStatsRequest>, I>>(object: I): QueryGetFeeStatsRequest {
    const message = createBaseQueryGetFeeStatsRequest();
    message.date = object.date ?? "";
    return message;
  },
};

function createBaseQueryGetFeeStatsResponse(): QueryGetFeeStatsResponse {
  return { feeStats: undefined };
}

export const QueryGetFeeStatsResponse = {
  encode(message: QueryGetFeeStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.feeStats !== undefined) {
      FeeStats.encode(message.feeStats, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetFeeStatsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetFeeStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.feeStats = FeeStats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeStatsResponse {
    return { feeStats: isSet(object.feeStats) ? FeeStats.fromJSON(object.feeStats) : undefined };
  },

  toJSON(message: QueryGetFeeStatsResponse): unknown {
    const obj: any = {};
    if (message.feeStats !== undefined) {
      obj.feeStats = FeeStats.toJSON(message.feeStats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetFeeStatsResponse>, I>>(base?: I): QueryGetFeeStatsResponse {
    return QueryGetFeeStatsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetFeeStatsResponse>, I>>(object: I): QueryGetFeeStatsResponse {
    const message = createBaseQueryGetFeeStatsResponse();
    message.feeStats = (object.feeStats !== undefined && object.feeStats !== null)
      ? FeeStats.fromPartial(object.feeStats)
      : undefined;
    return message;
  },
};

function createBaseQueryAllFeeStatsRequest(): QueryAllFeeStatsRequest {
  return { pagination: undefined };
}

export const QueryAllFeeStatsRequest = {
  encode(message: QueryAllFeeStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllFeeStatsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllFeeStatsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeStatsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllFeeStatsRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllFeeStatsRequest>, I>>(base?: I): QueryAllFeeStatsRequest {
    return QueryAllFeeStatsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllFeeStatsRequest>, I>>(object: I): QueryAllFeeStatsRequest {
    const message = createBaseQueryAllFeeStatsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllFeeStatsResponse(): QueryAllFeeStatsResponse {
  return { feeStats: [], pagination: undefined };
}

export const QueryAllFeeStatsResponse = {
  encode(message: QueryAllFeeStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.feeStats) {
      FeeStats.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllFeeStatsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllFeeStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.feeStats.push(FeeStats.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllFeeStatsResponse {
    return {
      feeStats: Array.isArray(object?.feeStats) ? object.feeStats.map((e: any) => FeeStats.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllFeeStatsResponse): unknown {
    const obj: any = {};
    if (message.feeStats?.length) {
      obj.feeStats = message.feeStats.map((e) => FeeStats.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllFeeStatsResponse>, I>>(base?: I): QueryAllFeeStatsResponse {
    return QueryAllFeeStatsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllFeeStatsResponse>, I>>(object: I): QueryAllFeeStatsResponse {
    const message = createBaseQueryAllFeeStatsResponse();
    message.feeStats = object.feeStats?.map((e) => FeeStats.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetFeeStatsByIndexesRequest(): QueryGetFeeStatsByIndexesRequest {
  return { startIndex: "", endIndex: "", pagination: undefined };
}

export const QueryGetFeeStatsByIndexesRequest = {
  encode(message: QueryGetFeeStatsByIndexesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.startIndex !== "") {
      writer.uint32(10).string(message.startIndex);
    }
    if (message.endIndex !== "") {
      writer.uint32(18).string(message.endIndex);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetFeeStatsByIndexesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetFeeStatsByIndexesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.startIndex = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.endIndex = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeStatsByIndexesRequest {
    return {
      startIndex: isSet(object.startIndex) ? String(object.startIndex) : "",
      endIndex: isSet(object.endIndex) ? String(object.endIndex) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryGetFeeStatsByIndexesRequest): unknown {
    const obj: any = {};
    if (message.startIndex !== "") {
      obj.startIndex = message.startIndex;
    }
    if (message.endIndex !== "") {
      obj.endIndex = message.endIndex;
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetFeeStatsByIndexesRequest>, I>>(
    base?: I,
  ): QueryGetFeeStatsByIndexesRequest {
    return QueryGetFeeStatsByIndexesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetFeeStatsByIndexesRequest>, I>>(
    object: I,
  ): QueryGetFeeStatsByIndexesRequest {
    const message = createBaseQueryGetFeeStatsByIndexesRequest();
    message.startIndex = object.startIndex ?? "";
    message.endIndex = object.endIndex ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetFeeStatsByDateRequest(): QueryGetFeeStatsByDateRequest {
  return { startDate: "", endDate: "", pagination: undefined };
}

export const QueryGetFeeStatsByDateRequest = {
  encode(message: QueryGetFeeStatsByDateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.startDate !== "") {
      writer.uint32(10).string(message.startDate);
    }
    if (message.endDate !== "") {
      writer.uint32(18).string(message.endDate);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetFeeStatsByDateRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetFeeStatsByDateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.startDate = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.endDate = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetFeeStatsByDateRequest {
    return {
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryGetFeeStatsByDateRequest): unknown {
    const obj: any = {};
    if (message.startDate !== "") {
      obj.startDate = message.startDate;
    }
    if (message.endDate !== "") {
      obj.endDate = message.endDate;
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetFeeStatsByDateRequest>, I>>(base?: I): QueryGetFeeStatsByDateRequest {
    return QueryGetFeeStatsByDateRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetFeeStatsByDateRequest>, I>>(
    object: I,
  ): QueryGetFeeStatsByDateRequest {
    const message = createBaseQueryGetFeeStatsByDateRequest();
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryUserStatsResponse(): QueryUserStatsResponse {
  return { stats: undefined };
}

export const QueryUserStatsResponse = {
  encode(message: QueryUserStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stats !== undefined) {
      UserStats.encode(message.stats, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserStatsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.stats = UserStats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryUserStatsResponse {
    return { stats: isSet(object.stats) ? UserStats.fromJSON(object.stats) : undefined };
  },

  toJSON(message: QueryUserStatsResponse): unknown {
    const obj: any = {};
    if (message.stats !== undefined) {
      obj.stats = UserStats.toJSON(message.stats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryUserStatsResponse>, I>>(base?: I): QueryUserStatsResponse {
    return QueryUserStatsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryUserStatsResponse>, I>>(object: I): QueryUserStatsResponse {
    const message = createBaseQueryUserStatsResponse();
    message.stats = (object.stats !== undefined && object.stats !== null)
      ? UserStats.fromPartial(object.stats)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  AssetStats(request: QueryDateRequest): Promise<QueryAssetStatsResponse>;
  UserStats(request: QueryDateRequest): Promise<QueryUserStatsResponse>;
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of FeeStats items. */
  FeeStats(request: QueryGetFeeStatsRequest): Promise<QueryGetFeeStatsResponse>;
  FeeStatsAll(request: QueryAllFeeStatsRequest): Promise<QueryAllFeeStatsResponse>;
}

export const QueryServiceName = "stwartchain.stats.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.AssetStats = this.AssetStats.bind(this);
    this.UserStats = this.UserStats.bind(this);
    this.Params = this.Params.bind(this);
    this.FeeStats = this.FeeStats.bind(this);
    this.FeeStatsAll = this.FeeStatsAll.bind(this);
  }
  AssetStats(request: QueryDateRequest): Promise<QueryAssetStatsResponse> {
    const data = QueryDateRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AssetStats", data);
    return promise.then((data) => QueryAssetStatsResponse.decode(_m0.Reader.create(data)));
  }

  UserStats(request: QueryDateRequest): Promise<QueryUserStatsResponse> {
    const data = QueryDateRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UserStats", data);
    return promise.then((data) => QueryUserStatsResponse.decode(_m0.Reader.create(data)));
  }

  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(_m0.Reader.create(data)));
  }

  FeeStats(request: QueryGetFeeStatsRequest): Promise<QueryGetFeeStatsResponse> {
    const data = QueryGetFeeStatsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeeStats", data);
    return promise.then((data) => QueryGetFeeStatsResponse.decode(_m0.Reader.create(data)));
  }

  FeeStatsAll(request: QueryAllFeeStatsRequest): Promise<QueryAllFeeStatsResponse> {
    const data = QueryAllFeeStatsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FeeStatsAll", data);
    return promise.then((data) => QueryAllFeeStatsResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
