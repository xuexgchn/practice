%%%-------------------------------------------------------------------
%%% @author xuexg
%%% @doc
%%%     ProjectEuler Problem 6
%%% @end
%%% Created : 17. 十月 2017 17:29
%%%-------------------------------------------------------------------
-module(p006).

%% API
-export([
    assert/0,
    solve/1
]).

assert() ->
    2640 = solve(10),
    25164150 = solve(100),
    test_ok.

solve(Num) ->
    SumA = lists:sum([X * X || X <- lists:seq(1, Num)]),
    SumB = lists:sum([X || X <- lists:seq(1, Num)]),
    SumB * SumB - SumA.

