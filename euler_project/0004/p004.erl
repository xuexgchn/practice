%%%-------------------------------------------------------------------
%%% @author xuexg
%%% @doc
%%%
%%% @end
%%% Created : 17. 十月 2017 15:42
%%%-------------------------------------------------------------------
-module(p004).

%% API
-export([solve/0]).

solve() ->
    PalindromeList = [X * Y ||
        X <- lists:seq(100, 999),
        Y <- lists:seq(100, 999),
        integer_to_list(X * Y) =:= lists:reverse(integer_to_list(X * Y))
        ],
    io:format("The answer is ~p~n", [lists:max(PalindromeList)]).

%% erlang shell run: p004:solve() = 906609