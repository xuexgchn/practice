%%%-------------------------------------------------------------------
%%% @author xuexg
%%% @doc
%%%     ProjectEuler Problem 2
%%% @end
%%% Created : 17. 十月 2017 14:10
%%%-------------------------------------------------------------------
-module(p002).

%% API
-export([solve/1]).

cal(Num, TmpAns, ReverseFibList) when hd(ReverseFibList) =< Num ->
    FirstFibNum = hd(ReverseFibList) + lists:nth(2, ReverseFibList),
    NewReVerseFibList = [FirstFibNum | ReverseFibList],
    Ans = case FirstFibNum rem 2 of
              0 -> TmpAns + FirstFibNum;
              1 -> TmpAns
          end,
    cal(Num, Ans, NewReVerseFibList);
cal(Num, Ans, ReverseFibList) when hd(ReverseFibList) > Num ->
    Ans.

solve(MaxNum) ->
    Ans = cal(MaxNum, 0, [1, 0]),
    io:format("The answer is ~p~n", [Ans]).

%% erlang shell run: p002:solve(4000000) = 4613732
    
