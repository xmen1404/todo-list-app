#!/bin/sh
sudo docker run --name server --net todo-list-app-net -p 8000:8000 xmen1404/todo-list-app-server