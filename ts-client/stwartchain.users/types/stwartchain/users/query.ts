/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import { Stats } from "./stats";
import { UniqueUsers } from "./unique_users";

export const protobufPackage = "stwartchain.users";

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

export interface QueryStatsByDateRequest {
  startDate: string;
  endDate: string;
  pagination: PageRequest | undefined;
}

export interface QueryStatsByDateResponse {
  stats: Stats[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUniqueUsersRequest {
  date: string;
}

export interface QueryGetUniqueUsersResponse {
  uniqueUsers: UniqueUsers | undefined;
}

export interface QueryAllUniqueUsersRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUniqueUsersResponse {
  uniqueUsers: UniqueUsers[];
  pagination: PageResponse | undefined;
}

export interface QueryTotalRequest {
}

export interface QueryTotalResponse {
  count: string;
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

function createBaseQueryStatsByDateRequest(): QueryStatsByDateRequest {
  return { startDate: "", endDate: "", pagination: undefined };
}

export const QueryStatsByDateRequest = {
  encode(message: QueryStatsByDateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryStatsByDateRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryStatsByDateRequest();
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

  fromJSON(object: any): QueryStatsByDateRequest {
    return {
      startDate: isSet(object.startDate) ? String(object.startDate) : "",
      endDate: isSet(object.endDate) ? String(object.endDate) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryStatsByDateRequest): unknown {
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

  create<I extends Exact<DeepPartial<QueryStatsByDateRequest>, I>>(base?: I): QueryStatsByDateRequest {
    return QueryStatsByDateRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryStatsByDateRequest>, I>>(object: I): QueryStatsByDateRequest {
    const message = createBaseQueryStatsByDateRequest();
    message.startDate = object.startDate ?? "";
    message.endDate = object.endDate ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryStatsByDateResponse(): QueryStatsByDateResponse {
  return { stats: [], pagination: undefined };
}

export const QueryStatsByDateResponse = {
  encode(message: QueryStatsByDateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.stats) {
      Stats.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryStatsByDateResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryStatsByDateResponse();
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

  fromJSON(object: any): QueryStatsByDateResponse {
    return {
      stats: Array.isArray(object?.stats) ? object.stats.map((e: any) => Stats.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryStatsByDateResponse): unknown {
    const obj: any = {};
    if (message.stats?.length) {
      obj.stats = message.stats.map((e) => Stats.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryStatsByDateResponse>, I>>(base?: I): QueryStatsByDateResponse {
    return QueryStatsByDateResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryStatsByDateResponse>, I>>(object: I): QueryStatsByDateResponse {
    const message = createBaseQueryStatsByDateResponse();
    message.stats = object.stats?.map((e) => Stats.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUniqueUsersRequest(): QueryGetUniqueUsersRequest {
  return { date: "" };
}

export const QueryGetUniqueUsersRequest = {
  encode(message: QueryGetUniqueUsersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.date !== "") {
      writer.uint32(10).string(message.date);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUniqueUsersRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUniqueUsersRequest();
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

  fromJSON(object: any): QueryGetUniqueUsersRequest {
    return { date: isSet(object.date) ? String(object.date) : "" };
  },

  toJSON(message: QueryGetUniqueUsersRequest): unknown {
    const obj: any = {};
    if (message.date !== "") {
      obj.date = message.date;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetUniqueUsersRequest>, I>>(base?: I): QueryGetUniqueUsersRequest {
    return QueryGetUniqueUsersRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetUniqueUsersRequest>, I>>(object: I): QueryGetUniqueUsersRequest {
    const message = createBaseQueryGetUniqueUsersRequest();
    message.date = object.date ?? "";
    return message;
  },
};

function createBaseQueryGetUniqueUsersResponse(): QueryGetUniqueUsersResponse {
  return { uniqueUsers: undefined };
}

export const QueryGetUniqueUsersResponse = {
  encode(message: QueryGetUniqueUsersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.uniqueUsers !== undefined) {
      UniqueUsers.encode(message.uniqueUsers, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUniqueUsersResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUniqueUsersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uniqueUsers = UniqueUsers.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryGetUniqueUsersResponse {
    return { uniqueUsers: isSet(object.uniqueUsers) ? UniqueUsers.fromJSON(object.uniqueUsers) : undefined };
  },

  toJSON(message: QueryGetUniqueUsersResponse): unknown {
    const obj: any = {};
    if (message.uniqueUsers !== undefined) {
      obj.uniqueUsers = UniqueUsers.toJSON(message.uniqueUsers);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryGetUniqueUsersResponse>, I>>(base?: I): QueryGetUniqueUsersResponse {
    return QueryGetUniqueUsersResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryGetUniqueUsersResponse>, I>>(object: I): QueryGetUniqueUsersResponse {
    const message = createBaseQueryGetUniqueUsersResponse();
    message.uniqueUsers = (object.uniqueUsers !== undefined && object.uniqueUsers !== null)
      ? UniqueUsers.fromPartial(object.uniqueUsers)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUniqueUsersRequest(): QueryAllUniqueUsersRequest {
  return { pagination: undefined };
}

export const QueryAllUniqueUsersRequest = {
  encode(message: QueryAllUniqueUsersRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUniqueUsersRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUniqueUsersRequest();
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

  fromJSON(object: any): QueryAllUniqueUsersRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUniqueUsersRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllUniqueUsersRequest>, I>>(base?: I): QueryAllUniqueUsersRequest {
    return QueryAllUniqueUsersRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllUniqueUsersRequest>, I>>(object: I): QueryAllUniqueUsersRequest {
    const message = createBaseQueryAllUniqueUsersRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUniqueUsersResponse(): QueryAllUniqueUsersResponse {
  return { uniqueUsers: [], pagination: undefined };
}

export const QueryAllUniqueUsersResponse = {
  encode(message: QueryAllUniqueUsersResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.uniqueUsers) {
      UniqueUsers.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUniqueUsersResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUniqueUsersResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.uniqueUsers.push(UniqueUsers.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllUniqueUsersResponse {
    return {
      uniqueUsers: Array.isArray(object?.uniqueUsers)
        ? object.uniqueUsers.map((e: any) => UniqueUsers.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUniqueUsersResponse): unknown {
    const obj: any = {};
    if (message.uniqueUsers?.length) {
      obj.uniqueUsers = message.uniqueUsers.map((e) => UniqueUsers.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllUniqueUsersResponse>, I>>(base?: I): QueryAllUniqueUsersResponse {
    return QueryAllUniqueUsersResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllUniqueUsersResponse>, I>>(object: I): QueryAllUniqueUsersResponse {
    const message = createBaseQueryAllUniqueUsersResponse();
    message.uniqueUsers = object.uniqueUsers?.map((e) => UniqueUsers.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryTotalRequest(): QueryTotalRequest {
  return {};
}

export const QueryTotalRequest = {
  encode(_: QueryTotalRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalRequest();
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

  fromJSON(_: any): QueryTotalRequest {
    return {};
  },

  toJSON(_: QueryTotalRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryTotalRequest>, I>>(base?: I): QueryTotalRequest {
    return QueryTotalRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryTotalRequest>, I>>(_: I): QueryTotalRequest {
    const message = createBaseQueryTotalRequest();
    return message;
  },
};

function createBaseQueryTotalResponse(): QueryTotalResponse {
  return { count: "" };
}

export const QueryTotalResponse = {
  encode(message: QueryTotalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.count !== "") {
      writer.uint32(10).string(message.count);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryTotalResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryTotalResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.count = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryTotalResponse {
    return { count: isSet(object.count) ? String(object.count) : "" };
  },

  toJSON(message: QueryTotalResponse): unknown {
    const obj: any = {};
    if (message.count !== "") {
      obj.count = message.count;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryTotalResponse>, I>>(base?: I): QueryTotalResponse {
    return QueryTotalResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryTotalResponse>, I>>(object: I): QueryTotalResponse {
    const message = createBaseQueryTotalResponse();
    message.count = object.count ?? "";
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
  /** Queries a list of StatsByDate items. */
  StatsByDate(request: QueryStatsByDateRequest): Promise<QueryStatsByDateResponse>;
  /** Queries a list of UniqueUsers items. */
  UniqueUsers(request: QueryGetUniqueUsersRequest): Promise<QueryGetUniqueUsersResponse>;
  UniqueUsersAll(request: QueryAllUniqueUsersRequest): Promise<QueryAllUniqueUsersResponse>;
  /** Queries a list of Total items. */
  Total(request: QueryTotalRequest): Promise<QueryTotalResponse>;
}

export const QueryServiceName = "stwartchain.users.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Stats = this.Stats.bind(this);
    this.StatsAll = this.StatsAll.bind(this);
    this.StatsByDate = this.StatsByDate.bind(this);
    this.UniqueUsers = this.UniqueUsers.bind(this);
    this.UniqueUsersAll = this.UniqueUsersAll.bind(this);
    this.Total = this.Total.bind(this);
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

  StatsByDate(request: QueryStatsByDateRequest): Promise<QueryStatsByDateResponse> {
    const data = QueryStatsByDateRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "StatsByDate", data);
    return promise.then((data) => QueryStatsByDateResponse.decode(_m0.Reader.create(data)));
  }

  UniqueUsers(request: QueryGetUniqueUsersRequest): Promise<QueryGetUniqueUsersResponse> {
    const data = QueryGetUniqueUsersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UniqueUsers", data);
    return promise.then((data) => QueryGetUniqueUsersResponse.decode(_m0.Reader.create(data)));
  }

  UniqueUsersAll(request: QueryAllUniqueUsersRequest): Promise<QueryAllUniqueUsersResponse> {
    const data = QueryAllUniqueUsersRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UniqueUsersAll", data);
    return promise.then((data) => QueryAllUniqueUsersResponse.decode(_m0.Reader.create(data)));
  }

  Total(request: QueryTotalRequest): Promise<QueryTotalResponse> {
    const data = QueryTotalRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Total", data);
    return promise.then((data) => QueryTotalResponse.decode(_m0.Reader.create(data)));
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
