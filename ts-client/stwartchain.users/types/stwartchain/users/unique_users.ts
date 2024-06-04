/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { UniqueUserAddresses } from "./unique_user_addresses";

export const protobufPackage = "stwartchain.users";

export interface UniqueUsers {
  date: string;
  uniqueUserAddresses: UniqueUserAddresses | undefined;
}

function createBaseUniqueUsers(): UniqueUsers {
  return { date: "", uniqueUserAddresses: undefined };
}

export const UniqueUsers = {
  encode(message: UniqueUsers, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.date !== "") {
      writer.uint32(10).string(message.date);
    }
    if (message.uniqueUserAddresses !== undefined) {
      UniqueUserAddresses.encode(message.uniqueUserAddresses, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UniqueUsers {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUniqueUsers();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.date = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.uniqueUserAddresses = UniqueUserAddresses.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UniqueUsers {
    return {
      date: isSet(object.date) ? String(object.date) : "",
      uniqueUserAddresses: isSet(object.uniqueUserAddresses)
        ? UniqueUserAddresses.fromJSON(object.uniqueUserAddresses)
        : undefined,
    };
  },

  toJSON(message: UniqueUsers): unknown {
    const obj: any = {};
    if (message.date !== "") {
      obj.date = message.date;
    }
    if (message.uniqueUserAddresses !== undefined) {
      obj.uniqueUserAddresses = UniqueUserAddresses.toJSON(message.uniqueUserAddresses);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UniqueUsers>, I>>(base?: I): UniqueUsers {
    return UniqueUsers.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UniqueUsers>, I>>(object: I): UniqueUsers {
    const message = createBaseUniqueUsers();
    message.date = object.date ?? "";
    message.uniqueUserAddresses = (object.uniqueUserAddresses !== undefined && object.uniqueUserAddresses !== null)
      ? UniqueUserAddresses.fromPartial(object.uniqueUserAddresses)
      : undefined;
    return message;
  },
};

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
