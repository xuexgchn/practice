%%%-------------------------------------------------------------------
%%% @author xuexg
%%% @doc
%%%
%%% @end
%%% Created : 17. 十月 2017 16:08
%%%-------------------------------------------------------------------
-module(p005).

%% API
-export([solve/0]).

gcd(A, B) when A rem B =:= 0 ->
    B;
gcd(A, B) ->
    gcd(B, A rem B).

cal_lcm(Num) ->
    cal_lcm(1, Num).

cal_lcm(Ans, Num) when Num > 1 ->
    Tmp = Ans * Num div gcd(Ans, Num),
    cal_lcm(Tmp, Num - 1);
cal_lcm(Ans, Num) when Num =:= 1 ->
    Ans.

solve() ->
    Ans = cal_lcm(20),
    io:format("The answer is ~p~n", [Ans]).

%% erlang shell run: p004:solve() = 232792560