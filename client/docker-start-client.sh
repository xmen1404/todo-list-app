#!/bin/sh
sudo docker run --name client --net todo-list-app-net -p 3000:3000 xmen1404/todo-list-app-client