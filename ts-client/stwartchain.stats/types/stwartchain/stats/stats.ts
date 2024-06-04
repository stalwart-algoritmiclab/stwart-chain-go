/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { AssetDailyStats } from "./asset_stats";

export const protobufPackage = "stwartchain.stats";

export interface AssetStats {
  dailyStats: AssetDailyStats | undefined;
}

function createBaseAssetStats(): AssetStats {
  return { dailyStats: undefined };
}

export const AssetStats = {
  encode(message: AssetStats, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dailyStats !== undefined) {
      AssetDailyStats.encode(message.dailyStats, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AssetStats {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAssetStats();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.dailyStats = AssetDailyStats.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): AssetStats {
    return { dailyStats: isSet(object.dailyStats) ? AssetDailyStats.fromJSON(object.dailyStats) : undefined };
  },

  toJSON(message: AssetStats): unknown {
    const obj: any = {};
    if (message.dailyStats !== undefined) {
      obj.dailyStats = AssetDailyStats.toJSON(message.dailyStats);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<AssetStats>, I>>(base?: I): AssetStats {
    return AssetStats.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<AssetStats>, I>>(object: I): AssetStats {
    const message = createBaseAssetStats();
    message.dailyStats = (object.dailyStats !== undefined && object.dailyStats !== null)
      ? AssetDailyStats.fromPartial(object.dailyStats)
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
