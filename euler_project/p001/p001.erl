%%%-------------------------------------------------------------------
%%% @author xuexg
%%% @doc
%%%     ProjectEuler Problem 1
%%% @end
%%% Created : 17. 十月 2017 14:08
%%%-------------------------------------------------------------------
-module(p001).

%% API
-export([solve/0]).

solve() ->
    F = fun(X) ->
        X rem 3 =:= 0 orelse X rem 5 =:= 0
        end,
    MultiplesList = [X || X <- lists:seq(1, 1000 - 1), F(X)],
    io:format("The answer is ~p~n", [lists:sum(MultiplesList)]).

%% erlang shell run: p001:solve() = 233168
