/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { Stats } from "./stats";

export const protobufPackage = "stwartchain.core";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetStatsRequest {
  date: string;
}

export interface QueryGetStatsResponse {
  stats: Stats | undefined;
}

export interface QueryAllStatsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllStatsResponse {
  stats: Stats[];
  pagination: PageResponse | undefined;
}

export interface QueryGetStatsByDateRequest {
  startDate: string;
  endDate: string;
  pagination: PageRequest | undefined;
}

export interface QueryModulesAddressesRequest {
}

export interface QueryModulesAddressesResponse {
  address: string[];
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

function createBaseQueryGetStatsRequest(): QueryGetStatsRequest {
  return { date: "" };
}

export const QueryGetStatsRequest = {
  encode(message: QueryGetStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.date !== "") {
      writer.uint32(10).string(message.date);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStatsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStatsRequest();
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

  fromJSON(object: any): QueryGetStatsRequest {
    return { date: isSet(object.date) ? String(object.date) : "" };
  },

  toJSON(message: QueryGetStatsRequest): unknown {
    const obj: any = {};
    if (message.date !== "") {
      obj.date = message.date;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetStatsRequest>, I>>(base?: I): QueryGetStatsRequest {
    return QueryGetStatsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetStatsRequest>, I>>(object: I): QueryGetStatsRequest {
    const message = createBaseQueryGetStatsRequest();
    message.date = object.date ?? "";
    return message;
  },
};

function createBaseQueryGetStatsResponse(): QueryGetStatsResponse {
  return { stats: undefined };
}

export const QueryGetStatsResponse = {
  encode(message: QueryGetStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stats !== undefined) {
      Stats.encode(message.stats, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStatsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.stats = Stats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetStatsResponse {
    return { stats: isSet(object.stats) ? Stats.fromJSON(object.stats) : undefined };
  },

  toJSON(message: QueryGetStatsResponse): unknown {
    const obj: any = {};
    if (message.stats !== undefined) {
      obj.stats = Stats.toJSON(message.stats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetStatsResponse>, I>>(base?: I): QueryGetStatsResponse {
    return QueryGetStatsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetStatsResponse>, I>>(object: I): QueryGetStatsResponse {
    const message = createBaseQueryGetStatsResponse();
    message.stats = (object.stats !== undefined && object.stats !== null) ? Stats.fromPartial(object.stats) : undefined;
    return message;
  },
};

function createBaseQueryAllStatsRequest(): QueryAllStatsRequest {
  return { pagination: undefined };
}

export const QueryAllStatsRequest = {
  encode(message: QueryAllStatsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStatsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStatsRequest();
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

  fromJSON(object: any): QueryAllStatsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllStatsRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllStatsRequest>, I>>(base?: I): QueryAllStatsRequest {
    return QueryAllStatsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllStatsRequest>, I>>(object: I): QueryAllStatsRequest {
    const message = createBaseQueryAllStatsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllStatsResponse(): QueryAllStatsResponse {
  return { stats: [], pagination: undefined };
}

export const QueryAllStatsResponse = {
  encode(message: QueryAllStatsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.stats) {
      Stats.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllStatsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllStatsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.stats.push(Stats.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllStatsResponse {
    return {
      stats: Array.isArray(object?.stats) ? object.stats.map((e: any) => Stats.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllStatsResponse): unknown {
    const obj: any = {};
    if (message.stats?.length) {
      obj.stats = message.stats.map((e) => Stats.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllStatsResponse>, I>>(base?: I): QueryAllStatsResponse {
    return QueryAllStatsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllStatsResponse>, I>>(object: I): QueryAllStatsResponse {
    const message = createBaseQueryAllStatsResponse();
    message.stats = object.stats?.map((e) => Stats.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetStatsByDateRequest(): QueryGetStatsByDateRequest {
  return { startDate: "", endDate: "", pagination: undefined };
}

export const QueryGetStatsByDateRequest = {
  encode(message: QueryGetStatsByDateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetStatsByDateRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetStatsByDateRequest();
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

  fromJSON(object: any): QueryGetStatsByDateRequest {
    return {
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryGetStatsByDateRequest): unknown {
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

  create<I extends Exact<DeepPartial<QueryGetStatsByDateRequest>, I>>(base?: I): QueryGetStatsByDateRequest {
    return QueryGetStatsByDateRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetStatsByDateRequest>, I>>(object: I): QueryGetStatsByDateRequest {
    const message = createBaseQueryGetStatsByDateRequest();
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryModulesAddressesRequest(): QueryModulesAddressesRequest {
  return {};
}

export const QueryModulesAddressesRequest = {
  encode(_: QueryModulesAddressesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryModulesAddressesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryModulesAddressesRequest();
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

  fromJSON(_: any): QueryModulesAddressesRequest {
    return {};
  },

  toJSON(_: QueryModulesAddressesRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryModulesAddressesRequest>, I>>(base?: I): QueryModulesAddressesRequest {
    return QueryModulesAddressesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryModulesAddressesRequest>, I>>(_: I): QueryModulesAddressesRequest {
    const message = createBaseQueryModulesAddressesRequest();
    return message;
  },
};

function createBaseQueryModulesAddressesResponse(): QueryModulesAddressesResponse {
  return { address: [] };
}

export const QueryModulesAddressesResponse = {
  encode(message: QueryModulesAddressesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.address) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryModulesAddressesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryModulesAddressesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.address.push(reader.string());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryModulesAddressesResponse {
    return { address: Array.isArray(object?.address) ? object.address.map((e: any) => String(e)) : [] };
  },

  toJSON(message: QueryModulesAddressesResponse): unknown {
    const obj: any = {};
    if (message.address?.length) {
      obj.address = message.address;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryModulesAddressesResponse>, I>>(base?: I): QueryModulesAddressesResponse {
    return QueryModulesAddressesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryModulesAddressesResponse>, I>>(
    object: I,
  ): QueryModulesAddressesResponse {
    const message = createBaseQueryModulesAddressesResponse();
    message.address = object.address?.map((e) => e) || [];
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Stats items. */
  Stats(request: QueryGetStatsRequest): Promise<QueryGetStatsResponse>;
  StatsAll(request: QueryAllStatsRequest): Promise<QueryAllStatsResponse>;
  /** Queries a list of ModulesAddresses items. */
  ModulesAddresses(request: QueryModulesAddressesRequest): Promise<QueryModulesAddressesResponse>;
}

export const QueryServiceName = "stwartchain.core.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Stats = this.Stats.bind(this);
    this.StatsAll = this.StatsAll.bind(this);
    this.ModulesAddresses = this.ModulesAddresses.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(_m0.Reader.create(data)));
  }

  Stats(request: QueryGetStatsRequest): Promise<QueryGetStatsResponse> {
    const data = QueryGetStatsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Stats", data);
    return promise.then((data) => QueryGetStatsResponse.decode(_m0.Reader.create(data)));
  }

  StatsAll(request: QueryAllStatsRequest): Promise<QueryAllStatsResponse> {
    const data = QueryAllStatsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "StatsAll", data);
    return promise.then((data) => QueryAllStatsResponse.decode(_m0.Reader.create(data)));
  }

  ModulesAddresses(request: QueryModulesAddressesRequest): Promise<QueryModulesAddressesResponse> {
    const data = QueryModulesAddressesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "ModulesAddresses", data);
    return promise.then((data) => QueryModulesAddressesResponse.decode(_m0.Reader.create(data)));
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
